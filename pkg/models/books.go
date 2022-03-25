package models

type Book struct {
	ID          string
	Name        string
	PageNumber  int
	StockNumber int
	Price       float64
	StockCode   string
	ISBN        string
	Author      Author
}

type Author struct {
	Name string
	ID   string
}
