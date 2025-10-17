package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

// Api for Get all Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please Sent GET Request", 400)
		return
	}
	utils.SendData(w, database.List(), 200)

}
