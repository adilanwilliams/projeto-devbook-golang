package main

import (
	"devbook/src/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Bootstrap()
	router := router.CreateRoute()

	fmt.Printf("Server listining on port %d\n", config.Port)
	erro := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
	if erro != nil {
		panic(erro)
	}

	log.Fatal()
}
