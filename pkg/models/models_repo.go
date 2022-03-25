package models

import "gorm.io/gorm"

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) Migrations() {
	br.db.AutoMigrate(&Book{})
	br.db.AutoMigrate(&Author{})
}

func (br *BookRepository) InsertSampleData(list []Book) {
	for _, book := range list {
		br.db.Where(Book{ID: book.ID}).
			Attrs(Book{ID: book.ID, Name: book.Name, PageNumber: book.PageNumber, StockNumber: book.StockNumber, Price: book.Price, StockCode: book.StockCode, ISBN: book.ISBN, Author: book.Author}).
			FirstOrCreate(&book)
	}
}
