package storage

import (
	"fmt"
)

//Desenvolvimento
const (
	Host = "localhost"
	Port = "5434" /*docker:*/
	//Port     = "5432" /*local:*/
	User     = "postgres"
	Password = "floripa@123"
	Dbname   = "postgres"
)

//ConnectarBasePostgres responsável por abrir a conexão com o postgres
func ConnectarBasePostgres() string {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		Host, Port, User, Password, Dbname)
	return psqlInfo
}
