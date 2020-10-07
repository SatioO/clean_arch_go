package handler

import (
	"net/http"

	"github.com/satioO/togo/app/adapter/port/usecase"
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
func (u *PermissionHandler) ListPermissions(w http.ResponseWriter, r *http.Request) {}

// CreatePermission ...
func (u *PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {}
