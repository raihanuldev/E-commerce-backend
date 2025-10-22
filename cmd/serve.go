package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	middilwares := middleware.NewMiddleware(cnf)
	productHandler := product.NewHandler(middilwares, &productRepo)
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(productHandler, userHandler, *cnf)
	server.Start()
}
