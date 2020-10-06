package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/satioO/togo/app/user"

	"github.com/satioO/togo/utils"

	"github.com/satioO/togo/domain"

	"github.com/gorilla/mux"
)

// UserHandler ...
type UserHandler struct {
	usecase user.Usecase
}

// RegisterUserHandler handles all user routes
func RegisterUserHandler(t *mux.Router, usecase user.Usecase) {
	handler := UserHandler{
		usecase,
	}

	t.HandleFunc("/users", handler.FindUsers).Methods(http.MethodGet)
	t.HandleFunc("/users", handler.SaveUser).Methods(http.MethodPost)
}

// FindUsers ...
func (u *UserHandler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := u.usecase.FindUsers(ctx)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// SaveUser ...
func (u *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := u.usecase.SaveUser(ctx, &user)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "User account created successfully")
}
