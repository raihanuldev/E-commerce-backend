package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Product Struct
type Product struct {
	ID          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Price       float32 `db:"price" json:"price"`
	ImgUrl      string  `db:"imgurl" json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	Delete(productId int) (*Product, error)
	Update(product Product) (*Product, error)
	List() ([]*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

//constructor or contructor Method

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
	INSERT INTO products(
    	title,
		description,
		price,
		ImgUrl
	) 	VALUES (
    	$1,
		$2,
		$3,
		$4
) 
	RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *productRepo) List() ([]*Product, error) {
	var products []*Product
	query := `SELECT id, title, description, price, imgurl FROM products;`

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *productRepo) Get(productId int) (*Product, error) {
	var product Product
	query := `SELECT id, title, description, price, imgurl FROM products  where id=$1;`
	err := r.db.Get(&product, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil

}
func (r *productRepo) Delete(productId int) (*Product, error) {
	query := `DELETE FROM products where id=$1`
	_, err := r.db.Exec(query, productId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
	UPDATE products SET (
    		title=$1,
			description=$2,
			price=$3,
			ImgUrl=$4
    	) where id=$5
	`
	rows := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	err := rows.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}
