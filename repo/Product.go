package repo

import (
	"database/sql"
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

//constructor or contructor Method

func NewProductRepo(db *sqlx.DB) *productRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
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
func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	var products []*domain.Product
	offset := ((page - 1) * limit) + 1
	query := `SELECT
				 id, 
				 title, 
				 description, 
				 price, 
				 imgurl 
			FROM products 
			LIMIT $1 
			OFFSET $2;`

	err := r.db.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}
func (r *productRepo) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM products`

	var count int64
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *productRepo) Get(productId int) (*domain.Product, error) {
	var product domain.Product
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
func (r *productRepo) Delete(productId int) (*domain.Product, error) {
	query := `DELETE FROM products where id=$1`
	_, err := r.db.Exec(query, productId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
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
