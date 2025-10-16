package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()

	mux := http.NewServeMux()
	WrapedMux := manager.WrapMux(mux, middleware.Logger, middleware.Hudai, middleware.Cors_Preflight)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", WrapedMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}
