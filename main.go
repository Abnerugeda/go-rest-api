package main

import (
	"fmt"

	"github.com/Abnerugeda/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando servidor Rest com go")
	routes.HandleRequest()
}
