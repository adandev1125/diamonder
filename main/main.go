package main

import (
	"fmt"
	"main/main/config"

	"main/main/routes"
	"net/http"
)

func main() {
	http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), routes.GetRouter())
}
