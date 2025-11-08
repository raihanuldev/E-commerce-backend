package cart

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newItem domain.CartItem
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newItem)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}
	//Database Push
	order, err := h.svc.AddToCart(newItem)
	if err != nil {
		http.Error(w, "Error to Add Cart", http.StatusInternalServerError)
	}
	utils.SendData(w, order, 201)

}
