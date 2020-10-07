package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	repository "github.com/satioO/togo/app/adapter/repository/mongo"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterRoleHandler() {
	repo := repository.NewRoleRepository(r.db, "roles")
	uc := usecase.NewRoleUsecase(repo)
	handle := handler.NewRoleHandler(uc)

	r.router.HandleFunc("/roles", handle.FindRoles).Methods(http.MethodGet)
	r.router.HandleFunc("/roles", handle.CreateRole).Methods(http.MethodPost)
}
