package main

import (
	"fmt"
	"net/http"

	"github.com/dougadriquei/desafioneoway/controller"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I am very keen on %s!", r.URL.Path[1:])
	success, error := controller.ReadFileController()
	fmt.Fprintf(w, "O status da leitura é:%v ", success)
	if success {
		return fmt.Fprintf(w, "\nsucesso")
	}

	fmt.Fprintf(w, fmt.Sprint(error))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
