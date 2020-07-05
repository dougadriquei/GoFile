package main

import (
	"encoding/json"
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
	fmt.Println("Passou 1")
	pathFile := "test/base_teste.txt"
	count, error := controller.ReadFileController(pathFile)
	data := result{
		QuantityInserted: count,
		Error:            error,
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8191", nil)
}
