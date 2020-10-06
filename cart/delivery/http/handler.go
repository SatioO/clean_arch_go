package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satioO/togo/domain"
)

type cartHandler struct{}

// RegisterCartHandler ...
func RegisterCartHandler(r *mux.Router) {
	handler := cartHandler{}

	r.HandleFunc("/cart", handler.GetCartItems)
}

// GetCartItems ...
func (c *cartHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
	// create a cart type
	cart := domain.Cart{}

	cart.Add([]domain.CartItem{
		domain.NextCoupan{ID: 1, Code: "Next", Discount: 10.0},
		domain.Product{ID: 1, Name: "iWatch", Price: 10.0},
		domain.GeneralCoupan{ID: 1, Code: "General", Discount: 10},
		domain.Product{ID: 2, Name: "iPhone", Price: 20.0},
	})

	fmt.Printf("Total Price: %.2f\n", cart.TotalPrice())

	response, _ := json.Marshal(map[string]map[string]domain.Product{"products": cart.Products})
	w.Write(response)
}
