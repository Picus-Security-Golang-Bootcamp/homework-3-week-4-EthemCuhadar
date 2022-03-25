package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string
	Name        string
	PageNumber  int
	StockNumber int
	Price       float64
	StockCode   string
	ISBN        string
	AuthorName  string
	Author      Author `gorm:"foreignKey:AuthorName"`
}

type Author struct {
	gorm.Model
	Name string
	ID   string
}

type BookList []Book

func (Book) TableName() string {
	return "Book"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %s, Name: %s, PageNumber: %d, StockNumber: %d, Price: %v, StockCode: %s, ISBN: %s, AuthorName: %s, AuthorID: %s, CreatedAt: %s",
		b.ID, b.Name, b.PageNumber, b.StockNumber, b.Price, b.StockCode, b.ISBN, b.Author.Name, b.Author.ID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Book (%s) deleting...", b.Name)
	return nil
}

func (b *Book) AfterDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s deleted", b.Name)
	return nil
}
