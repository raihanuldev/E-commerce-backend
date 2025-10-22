package product

import (
	"ecommerce/utils"
	"net/http"
)

// Api for Get all Products
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please Sent GET Request", 400)
		return
	}

	utils.SendData(w, h.productRepo.List(), 200)

}
