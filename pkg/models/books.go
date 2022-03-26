package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Book is struct type to store information about
// books. Author name field is foreign key for
// Author struct.
type Book struct {
	gorm.Model
	ID          string `gorm:"primary_key; unique"`
	Name        string `gorm:"size:50;not null"`
	PageNumber  int
	StockNumber int
	Price       float64
	StockCode   string
	ISBN        string
	AuthorName  string
	Author      *Author `gorm:"foreignKey:AuthorName"`
}

// Author struct to store name and author id
// for book authors.
type Author struct {
	gorm.Model
	Name string `gorm:"primary_key;unique"`
	ID   string
}

// BookList is a simple slice to store books
// that are read from CSV file and embedded
// into Book struct.
type BookList []Book

// TableName returns the table title for database.
func (Book) TableName() string {
	return "Book"
}

// ToString is string representation of relative
// fields for Book structs.
func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %s, Name: %s, PageNumber: %d, StockNumber: %d, Price: %v, StockCode: %s, ISBN: %s, AuthorName: %s, AuthorID: %s, CreatedAt: %s",
		b.ID, b.Name, b.PageNumber, b.StockNumber, b.Price, b.StockCode, b.ISBN, b.Author.Name, b.Author.ID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

// BeforeDelete returns nil and prints out information mesage
// just before deleting any item in database.
func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...", b.Name)
	return nil
}

// AfterDelete returns nil and prints out relative
// inforamtion just after deleting any item in database.
func (b *Book) AfterDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s deleted", b.Name)
	return nil
}
