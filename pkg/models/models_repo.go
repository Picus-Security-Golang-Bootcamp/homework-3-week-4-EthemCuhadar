package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// BookRepository stores a database via GORM.
type BookRepository struct {
	db *gorm.DB
}

// NewBookRepository simply creates and returns address
// of the database which is created.
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// Migrations automatically migrate book and author
// schemas, to keep the schema up to date.
func (br *BookRepository) Migrations() {
	bookPrototype := &Book{}
	authorPrototype := &Author{}
	br.db.AutoMigrate(bookPrototype, authorPrototype)
}

// InsertSampleData takes a book list and and insert the books
// and authors into schemas that are created in database.
func (br *BookRepository) InsertSampleData(list []Book) {
	for _, book := range list {
		br.db.Where(Book{ID: book.ID}).
			Attrs(Book{ID: book.ID, Name: book.Name}).
			FirstOrCreate(&book)
	}
}

////////////////////////////
// METHODS FOR BOOK MODEL //
////////////////////////////

//////////////////
// List Methods	//
//////////////////

// ListAllBooks finds and list out all items in
// book whose type is Book in the schema.
func (br *BookRepository) ListAllBooks() []Book {
	var books []Book
	br.db.Find(&books)
	return books
}

// ListAllBooksByAlphabeticOrder list out all the books
// in alphabetic order and returns them with an error if
// exists.
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

// ListBookByDescendingPriceOrder lists out all the book in the
// schema in terms of their prices in descending order.
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

// ListBookByAscendingPriceOrder lists out all the book in the
// schema in terms of their prices in ascending order.
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

// ListBookByAscendingStockNumberOrder lists out all the books
// in terms of their stock numbers in ascending order.
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

// ListBookByDescendingStockNumberOrder lists out all the books
// in terms of their stock numbers in descending order.
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

/////////////////
// Get Methods //
/////////////////

// GetBookByName Gets matched records for books in the database.
// Parameter does not have to be consistant of name totally. Function
// will return books with relates to name.
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

// GetBookByID gets matched records for books in the database.
func (br *BookRepository) GetBookByID(id string) (*Book, error) {
	var book Book
	result := br.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.ID)
	return &book, nil
}

// GetBookByAuthorName gets matched records for books in the database.
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

//GetBookByISBN gets matched records for books in the database.
func (br *BookRepository) GetBookByISBN(isbn string) (*Book, error) {
	var book Book
	result := br.db.Where(&Book{ISBN: isbn}).Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.ISBN)
	return &book, nil
}

// GetBookByMaxPriceLimit returns the books whose price is
// less than the price parameter.(e.g. Books whose price are less
// than 10 $.)
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

// GetBookByMinPriceLimit returns the books whose price is
// greater than the price parameter.(e.g. Books whose price are greater
// than 10 $.)
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

// GetBookWithMinPrice returns the book which has minimum
// price in the database.
func (br *BookRepository) GetBookWithMinPrice() (*Book, error) {
	var book Book
	result := br.db.Order("price asc").Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.Price)
	return &book, nil
}

// GetBookWithMaxPrice returns the book which has maximum
// price in the database.
func (br *BookRepository) GetBookWithMaxPrice() (*Book, error) {
	var book Book
	result := br.db.Order("price desc").Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println(book.Name, book.Price)
	return &book, nil
}

// GetBookWithPriceInterval returns the books which are between
// maximum and minimum prices which are queried. (e.g. Book whose
// price is between 10-15 $.)
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

// Create creates and stores new book in the database.
func (br *BookRepository) Create(book Book) error {
	result := br.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates and saves relative information in the database.
func (br *BookRepository) Update(book Book) error {
	result := br.db.Save(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete simply deletes the book from database.
func (br *BookRepository) Delete(book Book) error {
	result := br.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteBookByID deletes the book via its ID number from database.
func (br *BookRepository) DeleteBookByID(id string) error {
	result := br.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteBookByName deletes the book via its name from database.
func (br *BookRepository) DeleteBookByName(name string) error {
	result := br.db.Delete(&Book{}, name)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteBookByISBN deletes the book via its ISBN number from database.
func (br *BookRepository) DeleteBookByISBN(isbn string) error {
	result := br.db.Delete(&Book{}, isbn)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteBookByStockCode deletes the book via its Stock Code number from database.
func (br *BookRepository) DeleteBookByStockCode(sc string) error {
	result := br.db.Delete(&Book{}, sc)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//////////////////////////////
// METHODS FOR AUTHOR MODEL //
//////////////////////////////

/////////////////
// List Models //
/////////////////

// ListAllAuthors lists out all authors in the database.
func (br *BookRepository) ListAllAuthors() []Author {
	var authors []Author
	br.db.Find(&authors)
	for _, author := range authors {
		fmt.Println(author.Name)
	}
	return authors
}

// ListAllAuthorsByAlphabeticOrder lists out the authors in the
// database in alphabetic order.
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

////////////////
// Get Models //
////////////////

// GetBookNumberOfAutherByName lists out the book numbers of author whose
// name is queried.
func (br *BookRepository) GetBookNumberOfAutherByName(name string) (int64, error) {
	var count int64
	result := br.db.Model(&Book{}).Where("author_name = ?", name).Count(&count)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, result.Error
	}
	return count, nil
}
