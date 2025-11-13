package product

import (
	"ecommerce/utils"
	 _"ecommerce/domain"
	"net/http"
	"strconv"
	"strings"
)
// @Summary Delete a product
// @Description Delete a specific product by its ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "Product deleted successfully"
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Please send a DELETE request", http.StatusMethodNotAllowed)
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

	// Call repo to delete
	_, err = h.svc.Delete(pid)
	if err != nil {
		utils.SendData(w, map[string]string{"error": "Internal server error"}, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, map[string]string{"message": "Product deleted successfully"}, http.StatusOK)
}
