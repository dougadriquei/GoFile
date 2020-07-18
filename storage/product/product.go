package purchase

import (
	"database/sql"

	"GoFile/storage"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Plan []Plan `json:"plans"`
}

type Plan struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreatePurchases(product Product) (int, []error) {
	var error []error
	db, err := sql.Open("postgres", storage.ConnectarBasePostgres())
	if err != nil {
		error = append(error, err)
		return 0, error
	}
	//sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(1)
	//sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(1)
	defer db.Close()

	return 1, error
}
