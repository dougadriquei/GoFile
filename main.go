package main

import (
	"fmt"
	"net/http"

	"github.com/dougadriquei/desafioneoway/controller"
)

//TODO Melhorar handler, caso se transeformasse no padrão REST (token)
//TODO Pegar file da request
//TODO Utilizar variaveis de ambientes
//TODO Criar aquivos temporários? os.tempfile
//TODO melhorar imports do projeto. "github.com/

type result struct {
	QuantityInserted int     `json:"quantity_inserted,omitempty"`
	Error            []error `json:"error,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	count, error := controller.ReadFileController("test/base_teste.txt")
	fmt.Fprintf(w, "Count: %v \n", count)
	fmt.Fprintf(w, "Error: %v \n", fmt.Sprint(error))
	if error != nil {
		fmt.Fprintf(w, fmt.Sprint(error))
		return
	}
	data := result{
		QuantityInserted: count,
		Error:            error,
	}
	fmt.Fprintf(w, "A quantidade de registros inseridas é: %v\n", data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8095", nil)
}
