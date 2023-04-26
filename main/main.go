package main

import (
	"fmt"
	"main/main/config"

	"main/main/routes"
	"net/http"
)

func main() {
	fmt.Printf("Server started.\nYou can test on http://localhost:%d", config.PORT)

	http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), routes.GetRouter())
}
