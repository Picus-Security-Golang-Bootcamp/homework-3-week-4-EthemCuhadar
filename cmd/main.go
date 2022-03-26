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

	// bookRepo decleared
	bookRepo := models.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertSampleData(booklist)

	// Queries
	bookRepo.ListAllBooks()
	bookRepo.ListAllAuthorsByAlphabeticOrder()
	bookRepo.ListBookByDescendingPriceOrder()
	bookRepo.ListBookByAscendingPriceOrder()
	bookRepo.ListBookByDescendingStockNumberOrder()
	bookRepo.ListBookByAscendingStockNumberOrder()

	bookRepo.GetBookByName("Don Quiote")
	bookRepo.GetBookByID("1001")
	bookRepo.GetBookByAuthorName("Miguel De Cervantes Saavedra")
	bookRepo.GetBookByISBN("9142437239")
	bookRepo.GetBookByMaxPriceLimit(15.00)
	bookRepo.GetBookByMinPriceLimit(10.00)
	bookRepo.GetBookWithPriceInterval(10.00, 15.00)
	bookRepo.GetBookWithMaxPrice()
	bookRepo.GetBookWithMinPrice()

	sampleBook := models.Book{
		ID:          "1009",
		Name:        "The Brothers Karamazov",
		PageNumber:  569,
		StockNumber: 66,
		Price:       15.33,
		StockCode:   "362637",
		ISBN:        "1559963278",
		AuthorName:  "Fyodor Dostoevsky",
		Author:      &models.Author{Name: "Fyodor Dostoevsky", ID: "4593"},
	}
	bookRepo.Create(sampleBook)
	// Updated (price 15.33 -> 12.33)
	sampleBook.Price = 12.33
	bookRepo.Update(sampleBook)
	bookRepo.Delete(sampleBook)
	bookRepo.DeleteBookByID("1009")
	bookRepo.DeleteBookByName("The Brothers Karamazov")
	bookRepo.DeleteBookByISBN("1559963278")
	bookRepo.DeleteBookByStockCode("362637")

	bookRepo.ListAllAuthors()
	bookRepo.ListAllAuthorsByAlphabeticOrder()
	bookRepo.GetBookNumberOfAutherByName("Miguel De Cervantes Saavedra")
}
