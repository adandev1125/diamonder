package routes

import (
	"main/main/controllers"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var router *mux.Router
var once sync.Once

func GetRouter() *mux.Router {
	once.Do(func() {
		router = mux.NewRouter()
		initRouter()
	})

	return router
}

func initRouter() {

	router.HandleFunc("/", controllers.WelcomeHandler).Methods("GET")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
}
