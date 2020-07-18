package purchase

import (
	"GoFile/storage"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Plan []Plan `json:"plans"`
}

type Plan struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	IdProduct uint   `json:"idProduct"`
}

var DB *gorm.DB

func OpenDbConnection() (*gorm.DB, error) {
	//fmt.Println("Opening DB connection...")
	var err error
	DB, err = gorm.Open("postgres", storage.ConnectarBasePostgres())
	// if err != nil {
	// 	err = errors.New("Failed to connect to database")
	// 	return nil, err
	// }

	// Enable Logger, show detailed log
	DB.LogMode(false) // TODO Configurar por variável de ambiente

	configConnectionPool(DB)
	runMigrations(DB)

	return DB, err
}

func configConnectionPool(db *gorm.DB) {
	// Sets the maximum number of connections in the idle connection pool
	db.DB().SetMaxIdleConns(3)
	// Sets the maximum number of open connections to the database
	db.DB().SetMaxOpenConns(10)
	// Sets the maximum amount of time a connection may be reused
	db.DB().SetConnMaxLifetime(time.Hour)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&Product{})

	db.AutoMigrate(&Plan{})
	db.Model(&Plan{}).AddForeignKey("idproduct", "products(id)", "CASCADE", "CASCADE")

	// TODO @Douglas, precisa rodar migration pras tabelas novas inpc e formula?
}

func CreateProduct(product Product) (uint, error) {
	fmt.Println(product)
	OpenDbConnection()
	fmt.Println(DB)
	if db := DB.Create(&product); db.Error != nil {
		return 0, DB.Error
	}
	product.Plan[0].IdProduct = product.ID
	if db := DB.Create(&product.Plan[0]); db.Error != nil {
		return 0, DB.Error
	}
	product.Plan[1].IdProduct = product.ID
	if db := DB.Create(&product.Plan[1]); db.Error != nil {
		return 0, DB.Error
	}
	return product.ID, nil
}
