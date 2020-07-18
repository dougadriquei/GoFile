package main

import (
	"os"
	"os/signal"
	"syscall"

	"GoFile/http"
)

//TODO Melhorar handler, caso se transeformasse no padrão REST (token)
//TODO Pegar file da request
//TODO Utilizar variaveis de ambientes
//TODO Criar aquivos temporários? os.tempfile
//TODO melhorar imports do projeto. "github.com/

func main() {
	handler := http.NewHandler()
	server := http.New("8091", handler)
	server.ListenAndServe()
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}
