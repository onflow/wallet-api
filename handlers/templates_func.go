package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eqlabs/flow-wallet-service/templates"
	"github.com/gorilla/mux"
)

func (s *Templates) AddTokenFunc(rw http.ResponseWriter, r *http.Request) {
	var newToken templates.Token

	// Check body is not empty
	if err := checkNonEmptyBody(r); err != nil {
		handleError(rw, s.log, err)
		return
	}

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&newToken); err != nil {
		handleError(rw, s.log, InvalidBodyError)
		return
	}

	// Help marshalling the ID field if user for some reason tried to send a create request with an ID
	newToken.ID = 0

	// Add the token to database
	if err := s.service.AddToken(&newToken); err != nil {
		handleError(rw, s.log, err)
		return
	}

	handleJsonResponse(rw, http.StatusCreated, newToken)
}

func (s *Templates) ListTokensFunc(rw http.ResponseWriter, r *http.Request) {
	tokens, err := s.service.ListTokens()
	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	handleJsonResponse(rw, http.StatusOK, tokens)
}

func (s *Templates) GetTokenFunc(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	token, err := s.service.GetToken(id)
	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	handleJsonResponse(rw, http.StatusOK, token)
}

func (s *Templates) RemoveTokenFunc(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		handleError(rw, s.log, err)
		return
	}

	if err := s.service.RemoveToken(id); err != nil {
		handleError(rw, s.log, err)
		return
	}

	handleJsonResponse(rw, http.StatusOK, id)
}
