package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterPolicyHandler() {
	uc := usecase.NewPolicyUsecase()
	handle := handler.NewPolicyHandler(uc)

	r.router.HandleFunc("/policies", handle.ListPolicies).Methods(http.MethodGet)
	r.router.HandleFunc("/policies", handle.CreatePolicy).Methods(http.MethodPost)
}
