package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct { //step-1 for routes
	cnf            config.Config
	productHandler product.Handler
	userHandler    user.Handler
}

// step 2
func NewServer(productHandler *product.Handler, userHandler *user.Handler,cnf config.Config) *Server {
	return &Server{
		cnf:cnf,
		productHandler: *productHandler,
		userHandler:    *userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()
	// Type one
	manager.Use(middleware.PreFlight, middleware.Cors, middleware.Logger)

	// Type two
	// WrapedMux := manager.WrapMux(mux, middleware.Logger, middleware.PreFlight, middleware.Cors)
	WrapedMux := manager.WrapMux(mux)

	// initRoutes(mux, manager) //step3
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	port := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port, ", port)
	err := http.ListenAndServe(port, WrapedMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
		os.Exit(1)
	}
}
