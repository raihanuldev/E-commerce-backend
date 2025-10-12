package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GLobal Varibal
var productList []Product

// Struct
//
//	Product Struct
type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}

// Resubale Function
func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS,PATCH,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// APIS handle function

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Server apied created....")
}

// Api for Get all Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Please Sent GET Request", 400)
		return
	}
	sendData(w, productList, 200)

}

// Create Product
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	//cheking if request are vaild or invaild
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}
	//add product in global varibale
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(handleHello))
	mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))
	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))
	mux.Handle("OPTIONS /create-product", http.HandlerFunc(createProduct))

	//Server Config
	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}

// init function
func init() {
	// at frist i will create some product when init funtion excuted
	pr1 := Product{
		ID:          1,
		Title:       "headphone",
		Description: "This is Best Quality Headphone i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr2 := Product{
		ID:          2,
		Title:       "VIVO",
		Description: "This is Best Quality VIVO Phone i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr3 := Product{
		ID:          3,
		Title:       "APPLE",
		Description: "This is Best Quality Headphone i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr4 := Product{
		ID:          4,
		Title:       "SoundBox",
		Description: "This is Best Quality SoundBox i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr5 := Product{
		ID:          5,
		Title:       "Laptop",
		Description: "This is Best Quality Laptop i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr6 := Product{
		ID:          6,
		Title:       "Monitor",
		Description: "This is Best Quality Monitor i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	// for Static i will apend in slice
	productList = append(productList, pr1)
	productList = append(productList, pr2)
	productList = append(productList, pr3)
	productList = append(productList, pr4)
	productList = append(productList, pr5)
	productList = append(productList, pr6)
}
