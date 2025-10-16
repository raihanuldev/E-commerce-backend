package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()
	initRoutes(mux, manager)
	// mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	// mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	// // mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	// // mux.Handle("GET /products", middileware.Logger(http.HandlerFunc(handlers.GetProducts)))
	// // mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	// mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))
	globalRouterMux := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", globalRouterMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}
