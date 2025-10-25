package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please send a GET request", http.StatusBadRequest)
		return
	}

	products, err := h.productRepo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendData(w, products, http.StatusOK)
}
