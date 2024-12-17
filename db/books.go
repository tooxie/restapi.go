package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

// TODO: Pagination
func ListBooks() (*[]Book, error) {
	var books []Book
	res := Conn.Find(&books)
	if res.Error != nil {
		return nil, res.Error
	}

	return &books, nil
}

func CreateBook(book *Book) error {
	book.Slug = slug.Make(fmt.Sprintf("%s-%s", book.Author, book.Title))
	res := Conn.Create(&book)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func GetBook(id uuid.UUID, book *Book) error {
	res := Conn.Where("id = ?", id).First(&book)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func UpdateBook(book *Book) error {
	if book.Id == uuid.Nil {
		panic("Can't update book without ID")
	}
	book.Slug = slug.Make(fmt.Sprintf("%s-%s", book.Author, book.Title))
	res := Conn.Model(&book).Updates(book)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func DeleteBook(bookId string) error {
	var book Book
	err := Conn.Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return err
	}

	Conn.Delete(&book)
	return nil
}
