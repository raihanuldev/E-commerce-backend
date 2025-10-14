package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}",http.HandlerFunc(handlers.GetProductByID))
	globalRouterMux := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", globalRouterMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}
