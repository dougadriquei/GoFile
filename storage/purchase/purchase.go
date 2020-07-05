package purchase

import (
	"database/sql"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	brdoc "github.com/Nhanderu/brdoc"
	"github.com/dougadriquei/desafioneoway/storage"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Purchase Struct
type Purchase struct {
	CpfCnpj            string `json:"cpf_cnpj,omitempty"`
	Private            string `json:"private,omitempty"`
	Incompleted        string `json:"incompleted,omitempty"`
	LastPurchaseDate   string `json:"last_purchase_date,omitempty"`
	AverageTicket      string `json:"average_ticket,omitempty"`
	LastPurchaseTicket string `json:"last_purchase_ticket,omitempty"`
	MostFrequentStore  string `json:"most_frequent_store,omitempty"`
	LastPurchaseStore  string `json:"last_purchase_store,omitempty"`
}

//INSERT Constants das queries de criar e de trazer a quantidade de registros da entidade Purchase
//TODO - Acho que o CpfCnpj deveria ser uma das pks
const psqlInsert = `INSERT INTO purchase (cpf_cnpj, private, incompleted, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store)
VALUES  ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING idpurchase`

//INSERT query para persistir o modelo Purchase
const psqlCount = `SELECT COUNT(*) FROM purchase`

//isValidCpfOrCnpj valida cnpj ()
func isValidCpfOrCnpj(value string) bool {
	return brdoc.IsCPF(value) || brdoc.IsCNPJ(value)
}
func isValidCnpj(value string) bool {
	return brdoc.IsCNPJ(value)
}

//CreatePurchases persiste o model na entidade Purchase
func CreatePurchases(purchases []Purchase) (int, []error) {
	var error []error
	numberRecordsUnserted := 0
	fmt.Println("Passou 3")
	db, err := sql.Open("postgres", storage.ConnectarBasePostgres())
	if err != nil {
		error = append(error, err)
		return 0, error
	}
	fmt.Println("ERRO", err)
	fmt.Println("Passou 4")
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
			error = append(error, err)
			return 0, error
		}
		//Tratamento dos cpfs e cnpjs (remove a máscara, a partir do regex "[-,/,.]")
		cpfCnpj := regCpfCnpj.ReplaceAllString(purchase.CpfCnpj, "")
		mostFrequentStore := regCpfCnpj.ReplaceAllString(purchase.MostFrequentStore, "")
		lastPurchaseStore := regCpfCnpj.ReplaceAllString(purchase.LastPurchaseStore, "")
		lastPurchaseDate := purchase.LastPurchaseDate

		//Validação do cpfs e cnpjs
		if !isValidCpfOrCnpj(cpfCnpj) ||
			(mostFrequentStore != "NULL" && !isValidCnpj(mostFrequentStore)) ||
			(mostFrequentStore != "NULL" && !isValidCnpj(lastPurchaseStore)) {
			continue
		}

		private, _ := strconv.Atoi(purchase.Private)
		incompleted, _ := strconv.Atoi(purchase.Incompleted)
		averageTicket, _ := strconv.ParseFloat(strings.Replace(purchase.AverageTicket, ",", ".", 1), 32)
		lastPurchaseTicket, _ := strconv.ParseFloat(strings.Replace(purchase.LastPurchaseTicket, ",", ".", 1), 32)
		fmt.Println("Passou 5")
		//Executa a query do insert purchase, que possui uma ID como retorno.
		err = db.QueryRow(psqlInsert,
			//cpf_cnpj
			cpfCnpj,
			//private
			private,
			//incompleted
			incompleted,
			//last_purchase_date
			NewNullString(lastPurchaseDate),
			//average_ticket
			NewNullFloat64(math.Round(averageTicket*100)/100),
			//last_purchase_ticket
			NewNullFloat64(math.Round(lastPurchaseTicket*100)/100),
			//most_frequent_store
			NewNullString(mostFrequentStore),
			//last_purchase_store
			NewNullString(lastPurchaseStore),
		).Scan(&idpurchase)
		fmt.Println("Passou 6")
		if err != nil {
			fmt.Println("Passou 7")
			fmt.Println("ERRO", err)
			error = append(error, err)
			return 0, error
		}
		fmt.Println("Passou 8")

		idPurchase, err := strconv.Atoi(idpurchase)
		//Verifica se o registro foi persistido no postgres
		if err != nil || idPurchase == 0 {
			error = append(error, err)
		}
		numberRecordsUnserted++

	}
	if numberRecordsUnserted == 0 {
		return 0, error
	}
	return numberRecordsUnserted, error
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
