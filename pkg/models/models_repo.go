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

// METHODS FOR BOOK MODEL

// List Methods

func (br *BookRepository) ListAllBooks() []Book {
	var books []Book
	br.db.Find(&books)
	return books
}

func (br *BookRepository) ListAllBooksByAlphabeticOrder() ([]Book, error) {
	var books []Book
	result := br.db.Order("Name").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name)
	}
	return books, nil
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
	result := br.db.Order("Price").Order("Name").Find(&books)
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

// Get Methods

func (br *BookRepository) GetBookByName(bookName string) ([]Book, error) {
	var books []Book
	result := br.db.Where("name LIKE ?", "%"+bookName+"%").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name)
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
	result := br.db.Where("name LIKE ?", "%"+authorName+"%").Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name)
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

func (br *BookRepository) GetBookByMaxPriceLimit(pMax float64) ([]Book, error) {
	var books []Book
	result := br.db.Where("price < ?", pMax).Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.Price)
	}
	return books, nil
}

func (br *BookRepository) GetBookByMinPriceLimit(pMin float64) ([]Book, error) {
	var books []Book
	result := br.db.Where("price > ?", pMin).Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.Price)
	}
	return books, nil
}

func (br *BookRepository) GetBookWithMinPrice() (*Book, error) {
	var book Book
	result := br.db.Order("price asc").Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.Price)
	return &book, nil
}

func (br *BookRepository) GetBookWithMaxPrice() (*Book, error) {
	var book Book
	result := br.db.Order("price desc").Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.Price)
	return &book, nil
}

func (br *BookRepository) GetBookWithPriceInterval(pMin, pMax int) ([]Book, error) {
	var books []Book
	result := br.db.Where("price < ? AND price > ?", pMax, pMin).Find(&books)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range books {
		fmt.Println(book.Name, book.Price)
	}
	return books, nil
}

func (br *BookRepository) Create(book Book) error {
	result := br.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *BookRepository) Update(book Book) error {
	result := br.db.Save(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *BookRepository) Delete(book Book) error {
	result := br.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (br *BookRepository) DeleteBookById(id string) error {
	result := br.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (br *BookRepository) DeleteBookByName(name string) error {
	result := br.db.Delete(&Book{}, name)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (br *BookRepository) DeleteBookByISBN(isbn string) error {
	result := br.db.Delete(&Book{}, isbn)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (br *BookRepository) DeleteBookByStockCode(sc string) error {
	result := br.db.Delete(&Book{}, sc)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// METHODS FOR AUTHOR MODEL

// List Models

func (br *BookRepository) ListAllAuthors() []Author {
	var authors []Author
	br.db.Find(&authors)
	for _, author := range authors {
		fmt.Println(author.Name)
	}
	return authors
}

func (br *BookRepository) ListAllAuthorsByAlphabeticOrder() ([]Author, error) {
	var authors []Author
	result := br.db.Order("Name").Find(&authors)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	for _, book := range authors {
		fmt.Println(book.Name)
	}
	return authors, nil
}

func (br *BookRepository) GetBookNumberOfAutherByName(name string) (int64, error) {
	var count int64
	result := br.db.Model(&Book{}).Where("author_name = ?", name).Count(&count)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, result.Error
	}
	return count, nil
}
