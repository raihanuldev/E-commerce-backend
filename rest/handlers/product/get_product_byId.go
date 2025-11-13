package product

import (
	"ecommerce/utils"
	 _"ecommerce/domain"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Get product by ID
// @Description Retrieve a specific product using its unique ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} domain.Product
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
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

