package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/go-crud/dao"
	"github.com/lakshanwd/go-crud/repository"
)

//BookGetHandler - handle book get requests
func BookGetHandler(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	books, err := bookRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, convertListToArray(books))
}

//BookPostHandler - handle book post requests
func BookPostHandler(c *gin.Context) {
	bookRepo := repository.GetBookRepository()

	var book dao.Book
	if err := c.ShouldBindJSON(&book); err == nil {
		_, err = bookRepo.Insert(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//BookPutHandler - handle book put requests
func BookPutHandler(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	var book dao.Book
	if err := c.ShouldBindJSON(&book); err == nil {
		_, err = bookRepo.Update(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//BookDeleteHandler - handle book delete requests
func BookDeleteHandler(c *gin.Context) {
	bookRepo := repository.GetBookRepository()
	var book dao.Book
	book.BookID, _ = strconv.Atoi(c.Param("id"))
	_, err := bookRepo.Remove(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
