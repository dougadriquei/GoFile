package main

import (
	controller "desafioNeoway/controller"

	"fmt"
)

// TODO
//b) Docker Compose , com orientações para executar (arquivo readme)
//c) Boa organização lógica e documental (readme, comentários, etc);
//d) GIT

func main() {
	success, error := controller.ReadFileController()
	fmt.Print(success, error)
}
