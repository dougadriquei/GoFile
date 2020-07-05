package readfile

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	storage "github.com/dougadriquei/desafioneoway/storage/purchase"
)

//ReadFile responsável por ler o arquivo .txt e criar o DTO Purchase
func ReadFile(csvFile *os.File) ([]storage.Purchase, []error) {
	var error []error
	r := csv.NewReader(csvFile)
	//Separador utilizado na leitura = espaço em branco
	r.Comma = ' '
	purchases := make([]storage.Purchase, 0)

	//Verifica arquivo antes da leitura
	for {
		fileName := strings.Replace(csvFile.Name(), "resource/", "", -1)
		record, err := r.Read()
		if err == io.EOF {
			error = append(error, errors.New("Não foi possivel ler arquivo"))
			return purchases, error
		}
		if record[0] != "CPF" {
			error = append(error, errors.New("A versão do arquivo"+fileName+" ainda não foi implementada"))
			return purchases, error
		}
		break
	}
	//Percore todas as linhas do arquivo "base_teste.txt"
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if record[0] == "CPF" {
			continue
		}

		// Transforma array de cada linha em uma string.
		str := strings.Join(record, " ")
		// Captura da string todos os intervalos dos campos
		words := strings.Fields(str)
		//Preenche DTO
		cpfCnpj := words[0]
		private := words[1]
		incompleto := words[2]
		dataUltimaCompra := words[3]
		ticketMedio64 := words[4]
		ticketUltimaCompra64 := words[5]
		lojaMaisFrequente := words[6]
		lojaUltimaCompra := words[7]
		//Instancia DTO
		purchase := storage.Purchase{
			CpfCnpj:            cpfCnpj,
			Private:            private,
			Incompleted:        incompleto,
			LastPurchaseDate:   dataUltimaCompra,
			AverageTicket:      ticketMedio64,
			LastPurchaseTicket: ticketUltimaCompra64,
			MostFrequentStore:  lojaMaisFrequente,
			LastPurchaseStore:  lojaUltimaCompra,
		}
		//Adiciona DTO para um slice de Purchases
		purchases = append(purchases, purchase)
	}
	return purchases, error
}
