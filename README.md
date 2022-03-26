# Homework 3 | Week 4

----------------------------------------------

## Scope

This is a project in which a Bookstore database is built in Go programming language. In this project database server is chosen as PostgreSQL which is supported dialects in object-relational mapping, [GORM](gorm.io). 

-------------------------------------------------------------

## Requirements

* Golang

* GORM

* PostgreSQL

-------------------------------------------------------

## Usage

```[terminal]
go run main.go
```

----------------------------------------------------

## Queries

All of the query functions in the program are listed below as an example. Necessary query functions can be used to apply CRUD operations or searches.

* **ListAllBooks**()

* **ListAllAuthorsByAlphabeticOrder**()

* **ListBookByDescendingPriceOrder**()

* **ListBookByAscendingPriceOrder**()

* **ListBookByDescendingStockNumberOrder**()

* **ListBookByAscendingStockNumberOrder**()

* **GetBookByName**("*Don Quiote*")

* **GetBookByID**("*1001*")

* **GetBookByAuthorName**("*Miguel De Cervantes Saavedra*")

* **GetBookByISBN**("*9142437239*")

* **GetBookByMaxPriceLimit**(*15.00*)

* **GetBookByMinPriceLimit**(*10.00*)

* **GetBookWithPriceInterval**(*10.00*, *15.00*)

* **GetBookWithMaxPrice**()

* **GetBookWithMinPrice**()

* **Create**(&*sampleBook*)

* **Update**(&*sampleBook*)

* **Delete**(&*sampleBook*)

* **DeleteBookByID**("*1009*")

* **DeleteBookByName**("*The Brothers Karamazov*")

* **DeleteBookByISBN**("*1559963278*")

* **DeleteBookByStockCode**("*362637*")

* **ListAllAuthors**()

* **ListAllAuthorsByAlphabeticOrder**()

* **GetBookNumberOfAutherByName**("*Miguel De Cervantes Saavedra*")

---------------------------------------------------------

## Licence

MIT
