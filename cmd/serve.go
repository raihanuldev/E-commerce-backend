package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn,err:=db.NewConnection()

	if(err!=nil){
		fmt.Println(err)
		os.Exit(1)
	}


	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbConn)
	middilwares := middleware.NewMiddleware(cnf)
	productHandler := product.NewHandler(middilwares, &productRepo)
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(productHandler, userHandler, *cnf)
	server.Start()
}
