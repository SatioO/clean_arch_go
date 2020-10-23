package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	repository "github.com/satioO/togo/app/adapter/repository/mongo"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterResourceHandler() {
	repo := repository.NewResourceRepository(r.db, "resource")
	uc := usecase.NewResourceUsecase(repo)
	handle := handler.NewResourceHandler(uc)

	r.router.HandleFunc("/resources", handle.ListResources).Methods(http.MethodGet)
	r.router.HandleFunc("/resources", handle.ListResources).Methods(http.MethodPost)
}
