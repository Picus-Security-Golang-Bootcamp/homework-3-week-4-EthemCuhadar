package models

import "gorm.io/gorm"

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
