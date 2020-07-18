package controller

import (
	"errors"
	"fmt"

	domain "GoFile/readfile"
	daoStorage "GoFile/storage/product"
	dao "GoFile/storage/purchase"
	utils "GoFile/utils"
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

//CreateProduct controller do dominio
func CreateProduct(product daoStorage.Product) (uint, []error) {
	var error []error
	count, err := daoStorage.CreateProduct(product)
	if err != nil {
		error = append(error, err)
		return count, error
	}
	return count, error
}
