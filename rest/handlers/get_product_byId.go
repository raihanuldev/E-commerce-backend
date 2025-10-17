package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productid := r.PathValue("productId")
	fmt.Println(productid)
	pid, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "please give Valid ID", 400)
		return
	}
	for _, product := range database.ProductList {
		// log.Printf()
		if product.ID == pid {
			utils.SendData(w, product, 200)
			return
		}
	}
	utils.SendData(w, "Product Not found", 404)
}
