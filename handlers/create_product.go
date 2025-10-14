package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Create Product
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	//cheking if request are vaild or invaild
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}
	//add product in global varibale
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)

}
