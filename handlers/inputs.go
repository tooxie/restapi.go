package handlers

import (
	"github.com/google/uuid"
	"restapi/db"
)

type CreateBookInput struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateBookInput struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Slug   string
}

func (input *UpdateBookInput) Update(book *db.Book) {
	book.Id = input.Id
	book.Title = input.Title
	book.Author = input.Author
}
