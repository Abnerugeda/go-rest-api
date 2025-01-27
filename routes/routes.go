package routes

import (
	"log"
	"net/http"

	"github.com/Abnerugeda/go-rest-api/controllers"
	"github.com/Abnerugeda/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindOnePersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
