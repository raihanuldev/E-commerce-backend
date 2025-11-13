package order

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)
// @Summary Create a new order
// @Description Create a new order with product ID, user ID, quantity, and status
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body domain.Order true "Order data"
// @Success 201 {object} domain.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /order [post]

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newOrder domain.Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newOrder)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}
	//Database Push
	order, err := h.svc.CreateOrder(newOrder)
	if err != nil {
		http.Error(w, "Error to Place Order", http.StatusInternalServerError)
	}
	utils.SendData(w, order, 201)

}
