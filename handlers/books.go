package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"restapi/db"
	"restapi/errors"
)

func ListBooks(c *gin.Context) {
	books, err := db.ListBooks()
	if err != nil {
		log.Println(err)
		errors.Raise(c, errors.ErrorRetrievingBooks)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	log.Println("-- handlers.books.CreateBook()")
	var input CreateBookInput
	jsonErr := c.ShouldBindJSON(&input)
	if jsonErr != nil {
		errors.Raise(c, errors.InvalidInput)
		return
	}

	user := getCurrentlyLoggedInUser(c)
	book := db.Book{Title: input.Title, Author: input.Author, Owner: *user}
	err := db.CreateBook(&book)
	if err != nil {
		log.Println(err)
		errors.Raise(c, errors.ErrorSavingBook)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBook(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	var book db.Book
	err := db.GetBook(id, &book)
	if err != nil {
		errors.Raise(c, errors.BookNotFound)
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var input UpdateBookInput
	input.Id = uuid.MustParse(c.Param("id"))
	jsonErr := c.ShouldBindJSON(&input)
	if jsonErr != nil {
		log.Println("Error:", jsonErr.Error())
		errors.Raise(c, errors.InvalidInput)
		return
	}
	var book db.Book
	getErr := db.GetBook(input.Id, &book)
	if getErr != nil {
		log.Println(getErr)
		errors.Raise(c, errors.BookNotFound)
		return
	}

	input.Update(&book)
	updErr := db.UpdateBook(&book)
	if updErr != nil {
		log.Println(updErr)
		errors.Raise(c, errors.ErrorUpdatingBook)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book db.Book
	bookId := c.Param("id")
	err := db.DeleteBook(bookId)
	if err != nil {
		log.Println(err)
		errors.Raise(c, errors.BookNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
