package main

import (
	"fmt"
	"log"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init: ", err)
	}
	log.Println("Connected to local postgres server")
	if db == nil {
		fmt.Println("db is nil")
	}
}
