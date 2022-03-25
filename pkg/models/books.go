package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string `gorm:"primaryKey;unique"`
	Name        string `gorm:"size:50;not null"`
	PageNumber  int
	StockNumber int     `gorm:"type:INT;not null"`
	Price       float64 `gorm:"type:FLOAT8;not null"`
	StockCode   string  `gorm:"size:6;not null"`
	ISBN        string  `gorm:"size:10;not null"`
	Author      Author
}

type Author struct {
	gorm.Model
	Name string `gorm:"size:50;not null"`
	ID   string `gorm:"size:4;not null"`
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
