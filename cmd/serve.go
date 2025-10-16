package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()

	mux := http.NewServeMux()
	// Type one
	manager.Use(middleware.Cors, middleware.PreFlight, middleware.Logger)

	// Type two
	// WrapedMux := manager.WrapMux(mux, middleware.Logger, middleware.PreFlight, middleware.Cors)
	WrapedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", WrapedMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}
