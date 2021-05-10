// Package simple provides straightforward implementation for key management.
package simple

import (
	"context"
	"fmt"
	"log"

	"github.com/eqlabs/flow-wallet-service/data"
	"github.com/eqlabs/flow-wallet-service/keys"
	"github.com/eqlabs/flow-wallet-service/keys/encryption"
	"github.com/eqlabs/flow-wallet-service/keys/google"
	"github.com/eqlabs/flow-wallet-service/keys/local"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
)

type KeyManager struct {
	log             *log.Logger
	db              keys.KeyStore
	fc              *client.Client
	crypter         encryption.Crypter
	signAlgo        crypto.SignatureAlgorithm
	hashAlgo        crypto.HashAlgorithm
	adminAccountKey keys.Key
	cfg             Config
}

// NewKeyManager initiates a new key manager.
// It uses encryption.AESCrypter to encrypt and decrypt the keys.
func NewKeyManager(log *log.Logger, db keys.KeyStore, fc *client.Client) (result *KeyManager, err error) {
	cfg := ParseConfig()

	adminAccountKey := keys.Key{
		Index: cfg.AdminAccountKeyIndex,
		Type:  cfg.AdminAccountKeyType,
		Value: cfg.AdminAccountKeyValue,
	}

	crypter := encryption.NewAESCrypter([]byte(cfg.EncryptionKey))

	result = &KeyManager{
		log,
		db,
		fc,
		crypter,
		crypto.ECDSA_P256, // TODO: from config?
		crypto.SHA3_256,   // TODO: from config?
		adminAccountKey,
		cfg,
	}

	return
}

func (s *KeyManager) Generate(ctx context.Context, keyIndex, weight int) (result keys.Wrapped, err error) {
	switch s.cfg.DefaultKeyStorage {
	case keys.ACCOUNT_KEY_TYPE_LOCAL:
		result, err = local.Generate(
			s.signAlgo,
			s.hashAlgo,
			keyIndex,
			weight,
		)
	case keys.ACCOUNT_KEY_TYPE_GOOGLE_KMS:
		result, err = google.Generate(
			ctx,
			keyIndex,
			weight,
		)
	default:
		err = fmt.Errorf("keyStore.Generate() not implmented for %s", s.cfg.DefaultKeyStorage)
	}
	return
}

func (s *KeyManager) GenerateDefault(ctx context.Context) (keys.Wrapped, error) {
	return s.Generate(ctx, s.cfg.DefaultKeyIndex, s.cfg.DefaultKeyWeight)
}

func (s *KeyManager) Save(key keys.Key) (result data.Key, err error) {
	encValue, err := s.crypter.Encrypt([]byte(key.Value))
	if err != nil {
		return
	}
	result.Index = key.Index
	result.Type = key.Type
	result.Value = encValue
	return
}

func (s *KeyManager) Load(key data.Key) (result keys.Key, err error) {
	decValue, err := s.crypter.Decrypt([]byte(key.Value))
	if err != nil {
		return
	}
	result.Index = key.Index
	result.Type = key.Type
	result.Value = string(decValue)
	return
}

func (s *KeyManager) AdminAuthorizer(ctx context.Context) (keys.Authorizer, error) {
	return s.MakeAuthorizer(ctx, s.cfg.AdminAccountAddress)
}

func (s *KeyManager) UserAuthorizer(ctx context.Context, address string) (keys.Authorizer, error) {
	return s.MakeAuthorizer(ctx, address)
}

func (s *KeyManager) MakeAuthorizer(ctx context.Context, address string) (result keys.Authorizer, err error) {
	var key keys.Key

	result.Address = flow.HexToAddress(address)

	if address == s.cfg.AdminAccountAddress {
		key = s.adminAccountKey
	} else {
		var rawKey data.Key
		// Get the "least recently used" key for this address
		rawKey, err = s.db.AccountKey(address)
		if err != nil {
			return
		}
		key, err = s.Load(rawKey)
		if err != nil {
			return
		}
	}

	flowAcc, err := s.fc.GetAccount(ctx, flow.HexToAddress(address))
	if err != nil {
		return
	}

	result.Key = flowAcc.Keys[key.Index]

	var signer crypto.Signer

	// TODO: Decide whether we want to allow this kind of flexibility
	// or should we just panic if `key.Type` != `s.defaultKeyManager`
	switch key.Type {
	case keys.ACCOUNT_KEY_TYPE_LOCAL:
		signer, err = local.Signer(s.signAlgo, s.hashAlgo, key)
		if err != nil {
			break
		}
	case keys.ACCOUNT_KEY_TYPE_GOOGLE_KMS:
		signer, err = google.Signer(ctx, address, key)
		if err != nil {
			break
		}
	default:
		err = fmt.Errorf("key.Type not recognised: %s", key.Type)
	}

	result.Signer = signer

	return
}
