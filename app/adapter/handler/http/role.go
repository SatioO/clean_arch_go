package handler

import (
	"net/http"

	"github.com/satioO/togo/app/adapter/port/usecase"
)

// RoleHandler ...
type RoleHandler struct {
	uc usecase.RoleUsecase
}

// NewRoleHandler ...
func NewRoleHandler(uc usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{uc}
}

// FindRoles ...
func (u *RoleHandler) FindRoles(w http.ResponseWriter, r *http.Request) {}

// CreateRole ...
func (u *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {}
