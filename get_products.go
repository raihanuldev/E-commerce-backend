package main

import "net/http"

// Api for Get all Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please Sent GET Request", 400)
		return
	}
	sendData(w, productList, 200)

}
