package storage

import (
	"fmt"
)

//Desenvolvimento
const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "floripa@123"
	Dbname   = "postgres"
)

//ConnectarBasePostgres responsável por abrir a conexão com o postgres
func ConnectarBasePostgres() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)
	return psqlInfo
}
