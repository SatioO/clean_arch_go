package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	repository "github.com/satioO/togo/app/adapter/repository/mongo"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterPermissionHandler() {
	repo := repository.NewPermissionRepository(r.db, "permissions")
	uc := usecase.NewPermissionUsecase(repo)
	handle := handler.NewPermissionHandler(uc)

	r.router.HandleFunc("/permissions", handle.ListPermissions).Methods(http.MethodGet)
	r.router.HandleFunc("/permissions", handle.CreatePermission).Methods(http.MethodPost)
}
