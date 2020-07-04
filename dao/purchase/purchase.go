package purchase

import (
	"database/sql"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dougadriquei/desafioneoway/dao/storage"
	model "github.com/dougadriquei/desafioneoway/model/purchase"

	brdoc "github.com/Nhanderu/brdoc"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//INSERT query para persistir o modelo Purchase
const INSERT = `INSERT INTO purchase (cpf_cnpj, private, incompleted, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store)
VALUES  ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING idpurchase`

func isValidCpfOrCnpj(value string) bool {
	return brdoc.IsCPF(value) || brdoc.IsCNPJ(value)
}
func isValidCnpj(value string) bool {
	return brdoc.IsCNPJ(value)
}

//CreatePurchases persiste o model Purchase
func CreatePurchases(purchases []model.Purchase) (bool, []error) {
	var error []error
	db, err := sql.Open("postgres", storage.ConnectarBasePostgres())
	if err != nil {
		error = append(error, err)
	}
	//sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(1)
	//sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(1)
	defer db.Close()
	var idpurchase string
	for _, purchase := range purchases {
		// Make a Regex to say we only want numbers of "cpfs" and "cnpjs"
		regCpfCnpj, err := regexp.Compile("[-,/,.]")
		if err != nil {
			log.Fatal(err)
		}
		cpfCnpj := regCpfCnpj.ReplaceAllString(purchase.CpfCnpj, "")
		mostFrequentStore := regCpfCnpj.ReplaceAllString(purchase.MostFrequentStore, "")
		lastPurchaseStore := regCpfCnpj.ReplaceAllString(purchase.LastPurchaseStore, "")
		lastPurchaseDate := purchase.LastPurchaseDate

		//Validação do cpfs e cnps
		if !isValidCpfOrCnpj(cpfCnpj) ||
			(mostFrequentStore != "NULL" && !isValidCnpj(mostFrequentStore)) ||
			(mostFrequentStore != "NULL" && !isValidCnpj(lastPurchaseStore)) {
			continue
		}

		private, _ := strconv.Atoi(purchase.Private)
		incompleted, _ := strconv.Atoi(purchase.Incompleted)
		averageTicket, _ := strconv.ParseFloat(strings.Replace(purchase.AverageTicket, ",", ".", 1), 32)
		lastPurchaseTicket, _ := strconv.ParseFloat(strings.Replace(purchase.LastPurchaseTicket, ",", ".", 1), 32)
		//Executa a query do insert purchase, que possui uma ID como retorno.
		err = db.QueryRow(INSERT,
			//cpf_cnpj
			cpfCnpj,
			//private
			private,
			//incompleted
			incompleted,
			//last_purchase_date
			NewNullString(lastPurchaseDate),
			//average_ticket
			//TODO - Boa prática? Arredondamento do Ponto Flutuante
			NewNullFloat64(math.Round(averageTicket*100)/100),
			//last_purchase_ticket
			NewNullFloat64(math.Round(lastPurchaseTicket*100)/100),
			//most_frequent_store
			NewNullString(mostFrequentStore),
			//last_purchase_store
			NewNullString(lastPurchaseStore),
		).Scan(&idpurchase)

		if err != nil {
			error = append(error, err)
		}

		idPurchase, err := strconv.Atoi(idpurchase)
		//Verifica se o registro foi persistido no postgres
		if err != nil || idPurchase == 0 {
			error = append(error, err)
			return false, error
		}
	}

	return true, error
}

//NewNullString Verifica se a coluna veio nula da leitura e preenche a coluna como nil
func NewNullString(s string) sql.NullString {
	if len(s) == 0 || s == "NULL" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

//NewNullFloat64 Verifica se a coluna veio nula da leitura e preenche a coluna como nil
func NewNullFloat64(f float64) sql.NullFloat64 {
	if f == 0 {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: f,
		Valid:   true,
	}
}
