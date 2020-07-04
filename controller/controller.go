package controller

import (
	"errors"

	dao "github.com/dougadriquei/desafioneoway/dao/purchase"
	domain "github.com/dougadriquei/desafioneoway/domain/readfile"
	"github.com/dougadriquei/desafioneoway/util"
)

//ReadFileController controller do dominio
func ReadFileController() (bool, []error) {
	var error []error

	//Faz a leitura do arquivo: base_teste.txt
	csvfile := util.OpenFile("resource/base_teste.txt")
	defer csvfile.Close()

	//Varre cada linha do arquivo e captura um slice (array) das compras (purchases)
	purchases, err := domain.ReadFile(csvfile)
	if err != nil {
		return false, err
	}
	//Verifica se a lista está vazia
	if len(purchases) == 0 {
		error = append(error, errors.New("Lista está vazia"))
		return false, error
	}
	//Persiste todos os registros no Postgres
	success, error := dao.CreatePurchases(purchases)
	return success, error
}
