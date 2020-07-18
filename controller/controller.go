package controller

import (
	"errors"
	"fmt"

	utils "github.com/dougadriquei/GoFile/utils"

	domain "github.com/dougadriquei/GoFile/readfile"
	dao "github.com/dougadriquei/GoFile/storage/purchase"
)

//ReadFileController controller do dominio
func ReadFileController(pathFile string) (int, []error) {
	var error []error
	//Faz a leitura do arquivo: base_teste.txt
	csvfile := utils.OpenFile(pathFile)
	defer csvfile.Close()

	//Varre cada linha do arquivo e captura um slice (array) das compras (purchases)
	purchases, err := domain.ReadFile(csvfile)
	if err != nil {
		return 0, err
	}
	//Verifica se a lista está vazia
	if len(purchases) == 0 {
		error = append(error, errors.New("Lista está vazia"))
		return 0, error
	}
	//Persiste todos os registros no Postgres
	fmt.Println("Passou 2")
	count, err := dao.CreatePurchases(purchases)
	if err != nil {
		error = err
		return count, error
	}
	return count, error
}
