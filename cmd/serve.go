package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()
	// rest.Start(cnf)
	server := rest.NewServer(&product.Handler{}, &user.Handler{}, cnf)
	server.Start()
}
