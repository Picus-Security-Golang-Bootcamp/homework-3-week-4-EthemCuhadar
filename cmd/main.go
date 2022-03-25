package main

import (
	"fmt"
	"log"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/csv_utils"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/database"
)

var envFile = "../.env"
var csvFile = "../books.csv"

func main() {
	db, err := database.NewPsqlDB(envFile)
	if err != nil {
		log.Fatal("Postgres cannot init: ", err)
	}
	log.Println("Connected to local postgres server")
	if db == nil {
		fmt.Println("db is nil")
	}
	booklist, err := csv_utils.ReadCSV(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book Count: ", len(booklist))
}
