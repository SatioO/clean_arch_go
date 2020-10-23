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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resources, err := u.uc.ListResources(ctx)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resources)
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

	err := u.uc.CreateResource(ctx, resource)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Resource created successfully")
}
