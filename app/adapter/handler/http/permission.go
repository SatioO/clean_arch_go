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

// PermissionHandler ...
type PermissionHandler struct {
	uc usecase.PermissionUsecase
}

// NewPermissionHandler ...
func NewPermissionHandler(uc usecase.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{uc}
}

// ListPermissions ...
func (u *PermissionHandler) ListPermissions(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	permissions, err := u.uc.ListPermissions(ctx)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, permissions)
}

// CreatePermission ...
func (u *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	permission := domain.Permission{}
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := domain.ValidateAccessLevel(permission.AccessLevel)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = u.uc.CreatePermission(ctx, &permission)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Permission created successfully")
}
