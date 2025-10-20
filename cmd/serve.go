package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middilwares:= middleware.NewMiddleware(cnf)
	productHandler:= product.NewHandler(middilwares)
	userHandler:= user.NewHandler()
	

	server := rest.NewServer(productHandler,userHandler, *cnf)
	server.Start()
}
