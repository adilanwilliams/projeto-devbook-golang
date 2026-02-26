package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Port is the API port.
	Port = 0

	// ConnectionStringPGDatabase is the string connection with PostgreeSql database.
	ConnectionStringPGDatabase = ""
)

func Bootstrap(){
	var erro error

	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_POST"))
	if erro != nil{
		Port = 3030
	}

	ConnectionStringPGDatabase = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_DATABASE_HOST"),
		os.Getenv("PG_DATABASE_PORT"),
		os.Getenv("PG_DATABASE_USER"),
		os.Getenv("PG_DATABASE_PASSWORD"),
		os.Getenv("PG_DATABASE_NAME"),
	)
}