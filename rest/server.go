package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()

	mux := http.NewServeMux()
	// Type one
	manager.Use(middleware.PreFlight, middleware.Cors, middleware.Logger)

	// Type two
	// WrapedMux := manager.WrapMux(mux, middleware.Logger, middleware.PreFlight, middleware.Cors)
	WrapedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)
	port := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running on port, ", port)
	err := http.ListenAndServe(port, WrapedMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
		os.Exit(1)
	}
}
