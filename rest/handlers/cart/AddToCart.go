package cart

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GOT IT REQUEST")
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	usrID := 1
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

	order, err := h.svc.AddToCart(newItem, int64(usrID))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error to Add Cart", http.StatusInternalServerError)
	}
	utils.SendData(w, order, 201)

}
