package main

import (
	"devbook/src/router"
	"fmt"
	"net/http"
)

func main() {
	port := ":8000"
	router := router.CreateRoute()

	fmt.Printf("Listining on port %s", port)
	http.ListenAndServe(port, router)
}
