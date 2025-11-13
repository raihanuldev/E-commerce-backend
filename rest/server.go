package rest

import (
	"ecommerce/config"

	"ecommerce/rest/handlers/cart"
	"ecommerce/rest/handlers/order"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger"
    _ "ecommerce/docs"
)

type Server struct { //step-1 for routes
	cnf            config.Config
	productHandler product.Handler
	userHandler    user.Handler
	orderHandler   order.Handler
	cartHandler    cart.Handler
}

// step 2
func NewServer(productHandler *product.Handler, userHandler *user.Handler, orderHandler *order.Handler, cnf config.Config, cartHandler *cart.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: *productHandler,
		userHandler:    *userHandler,
		orderHandler:   *orderHandler,
		cartHandler:    *cartHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()
	// Type one
	manager.Use(middleware.PreFlight, middleware.Cors, middleware.Logger)

	//Swagger Documentaion
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Type two
	// WrapedMux := manager.WrapMux(mux, middleware.Logger, middleware.PreFlight, middleware.Cors)
	WrapedMux := manager.WrapMux(mux)

	// initRoutes(mux, manager) //step3
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.orderHandler.RegisterRoutes(mux, manager)
	server.cartHandler.RegisterRoutes(mux, manager)

	port := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port, ", port)
	err := http.ListenAndServe(port, WrapedMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
		os.Exit(1)
	}
}
