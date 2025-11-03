package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	productService "ecommerce/product"
	orderService "ecommerce/order"
	usrService "ecommerce/user"
	"ecommerce/repo"
	"ecommerce/rest"
	prodcutHandler "ecommerce/rest/handlers/product"
	usrhandler "ecommerce/rest/handlers/user"
	orderHandler "ecommerce/rest/handlers/order"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

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
	orderRepo:= repo.NewOrderRepo(dbConn)

	middilwares := middleware.NewMiddleware(cnf)

	//DOMAINS
	usrServices := usrService.NewService(userRepo)
	productServices := productService.NewService(productRepo)
	orderService:= orderService.NewService(orderRepo)

	//Handlers
	productHandler := prodcutHandler.NewHandler(middilwares, productServices)
	orderHandler:= orderHandler.NewHandler(middilwares,orderService)
	userHandler := usrhandler.NewHandler(usrServices)

	server := rest.NewServer(productHandler, userHandler,orderHandler, *cnf)
	server.Start()
}
