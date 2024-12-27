package main

import (
	"fmt"

	"github.com/Abnerugeda/go-rest-api/database"
	"github.com/Abnerugeda/go-rest-api/models"
	"github.com/Abnerugeda/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nOME 1", Historia: "HISOTIRA 1"},
		{Id: 2, Nome: "nOME 2", Historia: "HISOTIRA 2"},
	}

	database.ConnectDB()

	fmt.Println("Iniciando servidor Rest com go")
	routes.HandleRequest()
}
