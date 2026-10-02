package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Guiziin227/goCleanArc/internal/models"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("POST /users", h.addUser)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	users := h.usecases.GetAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h Handlers) addUser(w http.ResponseWriter, r *http.Request) {

	var req models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Invalid request body"})
		return
	}

	id, err := h.usecases.Add(req)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{NewUserID: id})
}
