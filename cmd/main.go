package main

import (
	"fmt"
	"log"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/csv_utils"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/database"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/models"
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
	// fmt.Println(booklist)
	fmt.Println("Book Count: ", len(booklist))
	bookRepo := models.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertSampleData(booklist)

	fmt.Println(bookRepo.ListAllBooks())
	bookRepo.GetBookByID("1003")
	fmt.Println(bookRepo.GetBookByAuthorName("J.R.R. Tolkein"))
	fmt.Println(bookRepo.GetBookByISBN("9142437239"))
}
