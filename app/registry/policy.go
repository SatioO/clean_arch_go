package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	repository "github.com/satioO/togo/app/adapter/repository/mongo"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterPolicyHandler() {
	repo := repository.NewPolicyRepository(r.db, "policy")
	uc := usecase.NewPolicyUsecase(repo)
	handle := handler.NewPolicyHandler(uc)

	r.router.HandleFunc("/policies", handle.ListPolicies).Methods(http.MethodGet)
	r.router.HandleFunc("/policies", handle.CreatePolicy).Methods(http.MethodPost)
}
