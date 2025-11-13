package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Summary Create a new product
// @Description Create a product and save it to the database
// @Tags Products
// @Accept json
// @Produce json
// @Param request body domain.Product true "Product data"
// @Success 201 {object} domain.Product
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /products [post]

// Create Product
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newProduct domain.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}

	// database.Store(newProduct
	pr, err := h.svc.Create(domain.Product(newProduct))
	if err != nil {
		http.Error(w, "Error to Create prodccut", http.StatusInternalServerError)
	}
	utils.SendData(w, pr, 201)

}
