package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Guiziin227/goCleanArc/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.usecases.GetAll(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

func (h *Handlers) getUserByID(w http.ResponseWriter, r *http.Request) {
	// Extraindo o ID do usuário da URL
	id := chi.URLParam(r, "id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.usecases.GetById(r.Context(), parsedID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

func (h *Handlers) addUser(w http.ResponseWriter, r *http.Request) {

	var req models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Invalid request body"})
		if err != nil {
			return
		}
		return
	}

	id, err := h.usecases.Add(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		err := json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(models.CreateUserResponse{NewUserID: id})
	if err != nil {
		return
	}
}
