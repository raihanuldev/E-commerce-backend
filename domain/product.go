package domain

//model or enitiy
type Product struct {
	ID          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Price       float32 `db:"price" json:"price"`
	ImgUrl      string  `db:"imgurl" json:"imgUrl"`
}
