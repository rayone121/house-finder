package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rayone121/reBank/backend/types"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("Method not allowed %s", r.Method)
	}

	var req types.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := s.store.GetAccountByUserName(req.UserName)
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("invalid password")
	}

	token, err := createJWT(acc)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
		Token:    token,
		Number:   acc.Number,
		UserName: acc.UserName,
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleTransferReq(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(types.TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}
	defer r.Body.Close()
	return writeJSON(w, http.StatusOK, transferReq)

}
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		accountID, err := getID(r)
		if err != nil {
			return err
		}
		account, err := s.store.GetAccountByID(accountID)
		if err != nil {
			return err
		}
		return writeJSON(w, http.StatusOK, account)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	account, err := types.NewAccount(req.FirstName, req.LastName, req.UserName, req.Password)
	if err != nil {
		return err
	}

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	accountID, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteAccount(accountID); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func getID(r *http.Request) (int, error) {
	id := mux.Vars(r)["id"]
	accountID, err := strconv.Atoi(id)
	if err != nil {
		return accountID, fmt.Errorf("invalid account ID: %s", id)
	}
	return accountID, nil
}
