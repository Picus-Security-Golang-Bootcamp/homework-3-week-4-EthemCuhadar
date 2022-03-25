package models

type Book struct {
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
	Name string `gorm:"size:50;not null"`
	ID   string `gorm:"size:4;not null"`
}
