package main

import (
	"encoding/json"
	"fmt"

	"github.com/Maxnetz/07-struct-interface/models"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(&obj, "", "\t")
	fmt.Println(string(raw))
}

type IBook interface {
	GetBook() *book
	Settitle(title string)
}

type book struct {
	*models.Book
}

func NewBook(title, author string) IBook {
	return &book{
		&models.Book{
			Title:  title,
			Author: author,
		}}
}

func (b *book) GetBook() *book {
	return b
}

func (b *book) Settitle(title string) {
	b.Title = title
}

func main() {
	// b := &book{
	// 	&models.Book{
	// 		Id:     1,
	// 		Title:  "Golang book",
	// 		Author: "Gopher",
	// 	},
	// }

	b := NewBook("The Hobbit", "J.R.R. Tolkien")
	b.Settitle("The Lord of the Rings")
	Debug(b.GetBook())
}
