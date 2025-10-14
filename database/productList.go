package database

// GLobal Varibal
var ProductList []Product

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
	ProductList = append(ProductList, pr1)
	ProductList = append(ProductList, pr2)
	ProductList = append(ProductList, pr3)
	ProductList = append(ProductList, pr4)
	ProductList = append(ProductList, pr5)
	ProductList = append(ProductList, pr6)
}
