package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)
// @Summary Update a product
// @Description Update a product’s details by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body domain.Product true "Updated product data"
// @Success 200 {object} domain.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please send a PUT request", http.StatusMethodNotAllowed)
		return
	}

	// Parse product ID from path
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 3 {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	productIDStr := parts[2]

	pid, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Please provide a valid ID", http.StatusBadRequest)
		return
	}

	// Parse request body JSON
	var updatedProduct domain.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updatedProduct.ID = pid // ensure ID is set

	// Call repo to update
	product, err := h.svc.Update(updatedProduct)
	if err != nil {
		utils.SendData(w, map[string]string{"error": "Internal server error"}, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
