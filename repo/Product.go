package repo

//	Product Struct
type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) Product
	Get(productId int) *Product
	Delete(productId int)
	Update(product Product)
	List() []Product
}

type productRepo struct {
	productList []Product
}

//constructor or contructor Method

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	GenerateInitProducts(repo)
	return repo
}

func (r *productRepo) List() []Product {
	return r.productList
}

func (r *productRepo) Create(p Product) Product {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, p)
	return p
}
func (r *productRepo) Get(productId int) *Product {
	for _, product := range r.productList {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}
func (r *productRepo) Delete(productId int) {
	var tmpList []Product
	for idx, product := range r.productList {
		if product.ID != productId {
			tmpList[idx] = product
		}
	}
	r.productList = tmpList
}
func (r *productRepo) Update(p Product) {
	for idx, product := range r.productList {
		// log.Printf()
		if product.ID == p.ID {
			r.productList[idx] = p
		}
	}
}

func GenerateInitProducts(r *productRepo) {
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
	r.productList = append(r.productList, pr1)
	r.productList = append(r.productList, pr2)
	r.productList = append(r.productList, pr3)
	r.productList = append(r.productList, pr4)
	r.productList = append(r.productList, pr5)
	r.productList = append(r.productList, pr6)
}
