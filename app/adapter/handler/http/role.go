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

// RoleHandler ...
type RoleHandler struct {
	uc usecase.RoleUsecase
}

// NewRoleHandler ...
func NewRoleHandler(uc usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{uc}
}

// ListRoles ...
func (u *RoleHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	roles, err := u.uc.ListRoles(ctx)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, roles)
}

// CreateRole ...
func (u *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	role := domain.Role{}
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := u.uc.CreateRole(ctx, &role)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "User role created successfully")
}
