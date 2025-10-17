package database

// GLobal Varibal
var productList []Product

func Store(p Product)Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func Get(productId int) *Product {
	for _, product := range productList {
		// log.Printf()
		if product.ID == productId {
			return &product
		}
	}
	return nil
}

func Update(p Product) {
	for idx, product := range productList {
		// log.Printf()
		if product.ID == p.ID {
			productList[idx] = p
		}
	}
}
func Delete(productId int) {
	var tmpList []Product
	for idx, product := range productList {
		// log.Printf()
		if product.ID != productId {
			tmpList[idx] = product
		}
	}
	productList = tmpList
}

func List() []Product {
	return productList
}

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
