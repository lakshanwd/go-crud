package handler

import (
	"net/http"
	"strconv"

	"../dao"
	"../repository"
	"github.com/gin-gonic/gin"
)

//BookGetHandler - handle book get requests
func BookGetHandler(c *gin.Context) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	books, err := bookRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, convertListToArray(books))
}

//BookPostHandler - handle book post requests
func BookPostHandler(c *gin.Context) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var book dao.Book
	if err = c.ShouldBindJSON(&book); err == nil {
		err = bookRepo.Insert(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//BookPutHandler - handle book put requests
func BookPutHandler(c *gin.Context) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var book dao.Book
	if err = c.ShouldBindJSON(&book); err == nil {
		err = bookRepo.Update(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//BookDeleteHandler - handle book delete requests
func BookDeleteHandler(c *gin.Context) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var book dao.Book
	book.BookID, _ = strconv.Atoi(c.Param("id"))
	err = bookRepo.Remove(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, true)
}
