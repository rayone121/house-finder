package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rayone121/house-finder/backend/types"
)

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return writeJSON(w, http.StatusOK, &types.Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
