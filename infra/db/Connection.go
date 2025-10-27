package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	conStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cnf.USER,
		cnf.Password,
		cnf.Host,
		cnf.PORT,
		cnf.Name,
	)
	if cnf.SslMode {
		conStr += " sslmode=enable"
	} else {
		conStr += " sslmode=disable"
	}
	return conStr
}
func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
