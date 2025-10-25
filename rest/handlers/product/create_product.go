package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Create Product
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newProduct repo.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}

	// database.Store(newProduct
	pr, err := h.productRepo.Create(repo.Product(newProduct))
	if err != nil {
		http.Error(w, "Error to Create prodccut", http.StatusInternalServerError)
	}
	utils.SendData(w, pr, 201)

}
