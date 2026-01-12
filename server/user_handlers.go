package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	commonsUser "github.com/LucasNav6/goauth/internal/commons/user"
	"github.com/LucasNav6/goauth/pkg/models"
)

// RegisterUserRoutes registers user management endpoints on the provided mux.
func RegisterUserRoutes(mux *http.ServeMux, config *models.Configuration, provider models.Provider) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreateUser(w, r, config, provider)
		case http.MethodGet:
			handleListUsers(w, r, config)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// path: /users/{id} or /users/by-email
		path := strings.TrimPrefix(r.URL.Path, "/users/")
		if strings.HasPrefix(path, "by-email") {
			// /users/by-email?email=...
			if r.Method != http.MethodGet {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			handleGetByEmail(w, r, config)
			return
		}

		id := path
		switch r.Method {
		case http.MethodGet:
			handleGetUser(w, r, config, id)
		case http.MethodPut:
			handleUpdateUser(w, r, config, id)
		case http.MethodDelete:
			handleDeleteUser(w, r, config, id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, config *models.Configuration, provider models.Provider) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req models.UserUnauthenticated
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use provider signup when available
	if provider != nil {
		// provider.SignUp expects *UserUnauthenticated
		created, err := provider.SignUp(config, &req)
		if err != nil {
			http.Error(w, fmt.Sprintf("signup error: %v", err), http.StatusInternalServerError)
			return
		}
		writeJSON(w, created)
		return
	}

	// Fallback: create user only
	created, err := commonsUser.Create(config, req)
	if err != nil {
		http.Error(w, fmt.Sprintf("create user error: %v", err), http.StatusInternalServerError)
		return
	}
	writeJSON(w, created)
}

func handleListUsers(w http.ResponseWriter, r *http.Request, config *models.Configuration) {
	users, err := commonsUser.GetAll(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, users)
}

func handleGetByEmail(w http.ResponseWriter, r *http.Request, config *models.Configuration) {
	q := r.URL.Query().Get("email")
	if q == "" {
		http.Error(w, "email query required", http.StatusBadRequest)
		return
	}
	u, err := commonsUser.GetByEmail(config, q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, u)
}

func handleGetUser(w http.ResponseWriter, r *http.Request, config *models.Configuration, id string) {
	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}
	u, err := commonsUser.GetByUUID(config, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, u)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, config *models.Configuration, id string) {
	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload struct {
		Name  *string `json:"name,omitempty"`
		Email *string `json:"email,omitempty"`
		Image *string `json:"image,omitempty"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if payload.Name == nil && payload.Email == nil && payload.Image == nil {
		http.Error(w, "nothing to update", http.StatusBadRequest)
		return
	}
	if payload.Email != nil {
		// basic validation delegated to commons
	}
	if err := commonsUser.Update(config, id, payload.Name, payload.Email, payload.Image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, config *models.Configuration, id string) {
	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}
	if err := commonsUser.Delete(config, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
