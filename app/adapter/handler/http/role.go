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

// CreateRoles ...
func (u *RoleHandler) CreateRoles(w http.ResponseWriter, r *http.Request) {}
