package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
	"strings"
)

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
	_, err = h.productRepo.Delete(pid)
	if err != nil {
		utils.SendData(w, map[string]string{"error": "Internal server error"}, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, map[string]string{"message": "Product deleted successfully"}, http.StatusOK)
}
