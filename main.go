package main

import (
	"fmt"
	"main/config"
	"main/routes"
	"main/system/database"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	fmt.Printf("Server is starting...\n\n")

	db := database.GetDatabase()
	defer db.Close()

	router := routes.GetRouter()

	// cors middleware
	methods := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	handler := handlers.CORS(methods, origins, headers)(router)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.PORT),
		Handler: handler,
	}

	s.ListenAndServe()
}
