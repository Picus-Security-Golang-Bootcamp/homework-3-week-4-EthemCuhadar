package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) Migrations() {
	bookPrototype := &Book{}
	authorPrototype := &Author{}
	br.db.AutoMigrate(bookPrototype, authorPrototype)
}

func (br *BookRepository) InsertSampleData(list []Book) {
	for _, book := range list {
		br.db.Where(Book{ID: book.ID}).
			Attrs(Book{ID: book.ID, Name: book.Name}).
			FirstOrCreate(&book)
	}
}

func (br *BookRepository) ListAllBooks() []Book {
	var books []Book
	br.db.Find(&books)
	return books
}

func (br *BookRepository) ListBookByDescendingPriceOrder() ([]Book, error) {
	var books []Book
	result := br.db.Order("Price desc").Order("Name").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.Price, "$")
	}
	return books, nil
}

func (br *BookRepository) ListBookByAscendingPriceOrder() ([]Book, error) {
	var books []Book
	result := br.db.Order("Price").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.Price, "$")
	}
	return books, nil
}

func (br *BookRepository) ListBookByAscendingStockNumberOrder() ([]Book, error) {
	var books []Book
	result := br.db.Order("stock_number").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.StockNumber)
	}
	return books, nil
}

func (br *BookRepository) ListBookByDescendingStockNumberOrder() ([]Book, error) {
	var books []Book
	result := br.db.Order("stock_number desc").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.StockNumber)
	}
	return books, nil
}

func (br *BookRepository) GetBookByID(id string) (*Book, error) {
	var book Book
	result := br.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.ID)
	return &book, nil
}

func (br *BookRepository) GetBookByAuthorName(authorName string) ([]Book, error) {
	var books []Book
	result := br.db.Where(&Book{AuthorName: authorName}).Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.Author)
	}
	return books, nil
}

func (br *BookRepository) GetBookByISBN(isbn string) (*Book, error) {
	var book Book
	result := br.db.Where(&Book{ISBN: isbn}).Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.ISBN)
	return &book, nil
}
