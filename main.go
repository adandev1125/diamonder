package main

import (
	"fmt"
	"main/routes"
	"main/system"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	fmt.Println("Server is starting...")

	// get configuration settings
	config := system.GetConfig()
	// initialize configuration
	config.Init()
	// close databases when function ends
	defer system.CloseDatabases()

	// get router
	router := routes.GetRouter()

	// set up cors middleware
	methods := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	handler := handlers.CORS(methods, origins, headers)(router)

	// create server with specified address and handler
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", system.GetConfig().Port),
		Handler: handler,
	}

	// start listening on specified port
	s.ListenAndServe()
}
