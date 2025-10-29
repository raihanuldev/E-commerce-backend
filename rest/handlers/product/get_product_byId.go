package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Split the path manually
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 { // e.g., ["", "products", "2"]
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	productIDStr := parts[2]
	fmt.Println("Product ID from path:", productIDStr)

	pid, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Please provide a valid ID", http.StatusBadRequest)
		return
	}

	product, err := h.svc.Get(pid)
	if err != nil {
		utils.SendData(w, map[string]string{"error": "Internal server error"}, http.StatusInternalServerError)
		return
	}

	if product == nil {
		utils.SendData(w, map[string]string{"error": "Product not found"}, http.StatusNotFound)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}

