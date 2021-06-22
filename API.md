# Wallet API REST HTTP Routes

The documents contains the REST routes provided by the Flow Wallet API.

These routes are also defined in the [OpenAPI specification for this service](openapi.yml).

## Functionality

### 1. Admin

- [x] Single admin account (hot wallet)
- [x] Create user accounts (using admin account)

### 2. Transaction Execution

- [x] Send an arbitrary transaction from the admin account
- [x] Send an arbitrary transaction from a user account

### 3. Fungible Tokens

- [x] Send fungible token withdrawals from admin account (FLOW, FUSD)
- [x] Detect fungible token deposits to admin account (FLOW, FUSD)
- [x] Send fungible token withdrawals from a user account (FLOW, FUSD)
- [x] Detect fungible token deposits to a user account (FLOW, FUSD)
- [x] View the fungible token balance of the admin account
- [x] View the fungible token balance of a user account

### 4. Non-Fungible Tokens

- [ ] Set up admin account with non-fungible token collections (`NFT.Collection`)
- [ ] Send non-fungible token withdrawals from admin account
- [ ] Detect non-fungible token deposits to admin account
- [ ] Set up a user account with non-fungible token collections (`NFT.Collection`)
- [ ] Send non-fungible token withdrawals from a user account
- [ ] Detect non-fungible token deposits to a user account
- [ ] View the non-fungible tokens owned by the admin account
- [ ] View the non-fungible tokens owned by a user account

---

## Accounts

### List all accounts

`GET /v1/accounts`

Example

```sh
curl --request GET \
  --url http://localhost:3000/v1/accounts
```

```json
[
  {
    "address": "0xf8d6e0586b0a20c7"
  },
  {
    "address": "0xe467b9dd11fa00df"
  }
]
```

---

### Get an account

`GET /v1/accounts/{address}`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")

Example

```sh
curl --request GET \
  --url http://localhost:3000/v1/accounts/0xf8d6e0586b0a20c7
```

```json
{
  "address": "0xf8d6e0586b0a20c7"
}
```

---

### Create an account

`POST /v1/accounts`

Example

```sh
curl --request POST \
  --url http://localhost:3000/v1/accounts
```

```json
{
  "address": "0xe467b9dd11fa00df"
}
```

---

## Transaction Execution

### Execute a transaction

> :warning: Not yet implemented

`POST /v1/accounts/{address}/transactions`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")

Body (JSON)

- `code`: The Cadence code to execute in the transaction
  - The code must always specify exactly one authorizer (i.e. `prepare(auth: AuthAccount)`)

Example

```sh
curl --request POST \
  --url http://localhost:3000/v1/accounts/0xf8d6e0586b0a20c7/transactions \
  --header 'Content-Type: application/json' \
  --data '{ "code": "transaction { prepare(auth: AuthAccount) { log(\"Hello, World!\") } }" }'
```

```json
{
  "transactionId": "18647b584a03345f3b2d2c4d9ab2c4179ae1b124a7f62ef9f33910e5ca8b353c",
  "error": null,
}
```

---

## Fungible Tokens

Supported tokens:
- `FLOW`
- `FUSD`

### List all tokens

`GET /v1/accounts/{address}/fungible-tokens`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")

Example

```sh
curl --request GET \
  --url http://localhost:3000/v1/accounts/0xf8d6e0586b0a20c7/fungible-tokens
```

```json
[
  {
    "name": "flow"
  },
  {
    "name": "fusd"
  }
]
```

---

### Get details of a token type

`GET /v1/accounts/{address}/fungible-tokens/{tokenName}`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the fungible token (e.g. "flow")

Example

```sh
curl --request GET \
  --url http://localhost:3000/v1/accounts/0xf8d6e0586b0a20c7/fungible-tokens/flow
```

```json
{
  "name": "flow", 
  "balance": "42.0"
}
```

---

### List all withdrawals of a token type

> :warning: Not yet implemented

`GET /v1/accounts/{address}/fungible-tokens/{tokenName}/withdrawals`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the fungible token (e.g. "flow")

---

### Get details of a token withdrawal

> :warning: Not yet implemented

`GET /v1/accounts/{address}/fungible-tokens/{tokenName}/withdrawals/{transactionId}`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the fungible token (e.g. "flow")
- `transactionId`: The Flow transaction ID for the withdrawal

---

### Create a token withdrawal

`POST /v1/accounts/{address}/fungible-tokens/{tokenName}/withdrawals`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the fungible token (e.g. "flow")

Body (JSON)

- `amount`: The number of tokens to transfer (e.g. "123.456")
  - Must be a fixed-point number with a maximum of 8 decimal places
- `recipient`: The Flow address of the recipient (e.g. "0xf8d6e0586b0a20c7")

Example

```sh
curl --request POST \
  --url http://localhost:3000/v1/accounts/0xf8d6e0586b0a20c7/fungible-tokens/fusd/withdrawls \
  --header 'Content-Type: application/json' \
  --data '{ "recipient": "0xe467b9dd11fa00df", "amount": "123.456" }'
```

```json
{
  "transactionId": "18647b584a03345f3b2d2c4d9ab2c4179ae1b124a7f62ef9f33910e5ca8b353c",
  "recipient": "0xe467b9dd11fa00df",
  "amount": "123.456"
}
```

---

## Non-Fungible Tokens

> :warning: Not yet implemented

### List all tokens

> :warning: Not yet implemented

`GET /v1/accounts/{address}/non-fungible-tokens`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")

Example

```sh
curl --request GET \
  --url http://localhost:3000/v1/accounts/0xf8d6e0586b0a20c7/non-fungible-tokens
```

---

### Get details of a token

> :warning: Not yet implemented

`GET /v1/accounts/{address}/non-fungible-tokens/{tokenName}`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the non-fungible token (e.g. "nba-top-shot-moment")

---

### List all withdrawals of a token type

> :warning: Not yet implemented

`GET /v1/accounts/{address}/non-fungible-tokens/{tokenName}/withdrawals`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the non-fungible token (e.g. "nba-top-shot-moment")

---

### Get details of a token withdrawal

> :warning: Not yet implemented

`GET /v1/accounts/{address}/non-fungible-tokens/{tokenName}/withdrawals/{transactionId}`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the non-fungible token (e.g. "nba-top-shot-moment")
- `transactionId`: The Flow transaction ID for the withdrawal

---

#### Create a token withdrawal

> :warning: Not yet implemented

`POST /v1/accounts/{address}/non-fungible-tokens/{tokenName}/withdrawals`

Parameters

- `address`: The address of the account (e.g. "0xf8d6e0586b0a20c7")
- `tokenName`: The name of the non-fungible token (e.g. "nba-top-shot-moment")

Body (JSON)

- `recipient`: The Flow address of the recipient (e.g. "0xf8d6e0586b0a20c7")