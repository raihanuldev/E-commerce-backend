package cmd

import (
	cartService "ecommerce/carts"
	"ecommerce/config"
	"ecommerce/infra/db"
	orderService "ecommerce/order"
	productService "ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	cartHandler "ecommerce/rest/handlers/cart"
	orderHandler "ecommerce/rest/handlers/order"
	prodcutHandler "ecommerce/rest/handlers/product"
	usrhandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	usrService "ecommerce/user"
	"fmt"
	"os"
)

// @title           My Go Backend API
// @version         1.0
// @description     API documentation for my backend project using net/http
// @termsOfService  http://swagger.io/terms/

// @contact.name    Raihanul Islam Sharif
// @contact.url     https://github.com/raihanuldev
// @contact.email   rihanulislam2015@gmail.com

// @host      localhost:8080
// @BasePath  /

func Serve() {
	cnf := config.GetConfig()

	dbConn, err := db.NewConnection(&cnf.DbConnection)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbConn, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Repos
	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)
	orderRepo := repo.NewOrderRepo(dbConn)
	cartRepo := repo.NewCartRepo(dbConn)

	middilwares := middleware.NewMiddleware(cnf)

	//DOMAINS
	usrServices := usrService.NewService(userRepo)
	productServices := productService.NewService(productRepo)
	orderService := orderService.NewService(orderRepo)
	cartService := cartService.NewService(cartRepo)
	//Handlers
	productHandler := prodcutHandler.NewHandler(middilwares, productServices)
	orderHandler := orderHandler.NewHandler(middilwares, orderService)
	userHandler := usrhandler.NewHandler(usrServices)
	cartHandler := cartHandler.NewHandler(middilwares, cartService)

	server := rest.NewServer(productHandler, userHandler, orderHandler, *cnf, cartHandler)
	server.Start()
}
