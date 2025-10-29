package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	productService "ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prodcutHandler "ecommerce/rest/handlers/product"
	usrhandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	usrService "ecommerce/user"
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

	middilwares := middleware.NewMiddleware(cnf)

	//DOMAINS
	usrServices := usrService.NewService(userRepo)
	productServices := productService.NewService(productRepo)

	productHandler := prodcutHandler.NewHandler(middilwares, productServices)
	userHandler := usrhandler.NewHandler(usrServices)

	server := rest.NewServer(productHandler, userHandler, *cnf)
	server.Start()
}
