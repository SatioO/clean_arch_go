package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/satioO/togo/app/adapter/port/usecase"
	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/utils"
)

// UserHandler ...
type UserHandler struct {
	uc usecase.UserUsecase
}

// NewUserHandler ...
func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

// FindUsers ...
func (u *UserHandler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := u.uc.FindUsers(ctx)

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

	err := u.uc.SaveUser(ctx, &user)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "User account created successfully")
}
