package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/satioO/togo/domain"

	"github.com/gorilla/mux"
)

type productHandler struct{}

// RegisterProductHandler registers product handlers
func RegisterProductHandler(r *mux.Router) {
	handler := &productHandler{}

	r.HandleFunc("/products", handler.GetProducts)
	r.HandleFunc("/products/{productID}", handler.GetProductsByID)
}

func (t *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	t1 := domain.Product{ID: 1, Name: "Golang"}
	t2 := domain.Product{ID: 2, Name: "Serverless"}
	t3 := domain.Product{ID: 3, Name: "AWS"}
	t4 := domain.Product{ID: 4, Name: "Kubernetes"}

	response, _ := json.Marshal([]domain.Product{t1, t2, t3, t4})
	w.Write(response)
}

func (t *productHandler) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	strconv.Atoi(mux.Vars(r)["productID"])

	w.Write([]byte("{productID: ''}"))
}
