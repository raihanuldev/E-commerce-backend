package order

import (
	"ecommerce/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

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
	
}
