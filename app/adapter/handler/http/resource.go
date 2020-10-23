package handler

import (
	"net/http"

	"github.com/satioO/togo/app/adapter/port/usecase"
)

// ResourceHandler ...
type ResourceHandler struct {
	uc usecase.ResourceUsecase
}

// NewResourceHandler ...
func NewResourceHandler(uc usecase.ResourceUsecase) *ResourceHandler {
	return &ResourceHandler{uc}
}

// ListResources ...
func (u *ResourceHandler) ListResources(w http.ResponseWriter, r *http.Request) {

}

// CreateResource ...
func (u *ResourceHandler) CreateResource(w http.ResponseWriter, r *http.Request) {

}
