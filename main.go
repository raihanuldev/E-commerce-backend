package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
	"fmt"
)

func main() {

	serviceName := config.GetConfig().ServiceName
	fmt.Println(serviceName)

	cmd.Serve()
}
