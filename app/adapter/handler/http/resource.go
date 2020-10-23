package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/satioO/togo/app/domain"
	"github.com/satioO/togo/utils"

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resource := domain.Resource{}
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.uc.CreateResource(ctx, resource)
}
