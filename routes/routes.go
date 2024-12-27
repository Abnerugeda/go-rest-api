package routes

import (
	"log"
	"net/http"

	"github.com/Abnerugeda/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
