package csv_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-EthemCuhadar/pkg/models"
)

// readCSV reads book information in a CSV file
// and returns a book slice to store the books
// into database.
func ReadCSV(filename string) (models.BookList, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var booklist models.BookList
	for _, line := range records[1:] {

		// pageNumber convert from string
		pageNumber, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("Page Number error")
		}
		// stockNumber convert from string
		stockNumber, err := strconv.Atoi(line[3])
		if err != nil {
			fmt.Println("Stock Number error")
		}
		// price convert from string
		price, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			fmt.Println("Stock Number error")
		}
		// author declared from csv file
		author := &models.Author{Name: line[7], ID: line[8]}

		booklist = append(booklist, models.Book{
			ID:          line[0],
			Name:        line[1],
			PageNumber:  pageNumber,
			StockNumber: stockNumber,
			Price:       price,
			StockCode:   line[5],
			ISBN:        line[6],
			AuthorName:  line[7],
			Author:      author,
		})
	}
	return booklist, nil
}
