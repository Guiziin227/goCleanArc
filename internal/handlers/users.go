package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/google/uuid"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("POST /users", h.addUser)
	http.HandleFunc("GET /users/{id}", h.getUserById)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.usecases.GetAll(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h Handlers) getUserById(w http.ResponseWriter, r *http.Request) {
	// Extraindo o ID do usuário da URL
	id := r.URL.Path[len("/users/"):]

	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.usecases.GetById(r.Context(), parsedID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h Handlers) addUser(w http.ResponseWriter, r *http.Request) {

	var req models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Invalid request body"})
		return
	}

	id, err := h.usecases.Add(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{NewUserID: id})
}
