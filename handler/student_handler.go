package handler

import (
	"net/http"
	"strconv"

	"../dao"
	"../repository"
	"github.com/gin-gonic/gin"
)

//StudentPostHandler - handle student post request
func StudentPostHandler(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	var student dao.Student
	if err := c.ShouldBindJSON(&student); err == nil {
		err = studentRepo.Insert(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//StudentGetHandler - handle student get request
func StudentGetHandler(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	students, err := studentRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, convertListToArray(students))
}

//StudentPutHandler - handle student put request
func StudentPutHandler(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	var student dao.Student
	if err := c.ShouldBindJSON(&student); err == nil {
		err = studentRepo.Update(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//StudentDeleteHandler - handle student delete request
func StudentDeleteHandler(c *gin.Context) {
	studentRepo := repository.GetStudentRepository()
	var student dao.Student
	student.StudentID, _ = strconv.Atoi(c.Param("id"))
	err := studentRepo.Remove(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
