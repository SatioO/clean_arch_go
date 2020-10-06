package registry

import (
	"net/http"

	handler "github.com/satioO/togo/app/adapter/handler/http"
	"github.com/satioO/togo/app/adapter/repository"
	"github.com/satioO/togo/app/usecase"
)

func (r *registry) RegisterUserHandler() {
	repo := repository.NewUserRepository(r.db, "users")
	uc := usecase.NewUserUsecase(repo)
	handle := handler.NewUserHandler(uc)

	r.router.HandleFunc("/users", handle.FindUsers).Methods(http.MethodGet)
	r.router.HandleFunc("/users", handle.SaveUser).Methods(http.MethodPost)
}
