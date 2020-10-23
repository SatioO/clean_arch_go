package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterResourceHandler() {
	uc := usecase.NewResourceUsecase()
	handle := handler.NewResourceHandler(uc)

	r.router.HandleFunc("/resources", handle.ListResources).Methods(http.MethodGet)
	r.router.HandleFunc("/resources", handle.CreateResource).Methods(http.MethodPost)
}
