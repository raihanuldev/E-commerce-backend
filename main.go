package main

import (
	"fmt"
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	mux.Handle("/products", http.HandlerFunc(getProducts))
	mux.Handle("/create-product", http.HandlerFunc(createProduct))

	//Call all cores releted
	globalRouterMux := globalRouter(mux)
	//Server Config
	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", globalRouterMux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}


// init function
func init() {
	// at frist i will create some product when init funtion excuted
	pr1 := Product{
		ID:          1,
		Title:       "headphone",
		Description: "This is Best Quality Headphone i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr2 := Product{
		ID:          2,
		Title:       "VIVO",
		Description: "This is Best Quality VIVO Phone i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr3 := Product{
		ID:          3,
		Title:       "APPLE",
		Description: "This is Best Quality Headphone i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr4 := Product{
		ID:          4,
		Title:       "SoundBox",
		Description: "This is Best Quality SoundBox i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr5 := Product{
		ID:          5,
		Title:       "Laptop",
		Description: "This is Best Quality Laptop i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	pr6 := Product{
		ID:          6,
		Title:       "Monitor",
		Description: "This is Best Quality Monitor i have ever seen",
		Price:       "199.21",
		ImgUrl:      "https/////image.url",
	}
	// for Static i will apend in slice
	productList = append(productList, pr1)
	productList = append(productList, pr2)
	productList = append(productList, pr3)
	productList = append(productList, pr4)
	productList = append(productList, pr5)
	productList = append(productList, pr6)
}
