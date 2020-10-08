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

// PolicyHandler ...
type PolicyHandler struct {
	uc usecase.PolicyUsecase
}

// NewPolicyHandler ...
func NewPolicyHandler(uc usecase.PolicyUsecase) *PolicyHandler {
	return &PolicyHandler{uc}
}

// ListPolicies ...
func (u *PolicyHandler) ListPolicies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Policys, err := u.uc.ListPolicies(ctx)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, Policys)
}

// CreatePolicy ...
func (u *PolicyHandler) CreatePolicy(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	policy := domain.Policy{}
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := u.uc.CreatePolicy(ctx, &policy)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Policy created successfully")
}
