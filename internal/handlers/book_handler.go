package handlers

import (
	"net/http"
	"strconv"

	"api-perpus-go/internal/models"
	"api-perpus-go/internal/services"

	"github.com/gin-gonic/gin"
)

type bookReq struct {
	Title      string `form:"title" json:"title"`
	Author     string `form:"author" json:"author"`
	Year       int    `form:"year" json:"year"`
	Pages      int    `form:"pages" json:"pages"`
	CategoryID *int   `form:"category_id" json:"category_id"`
}


func CreateBook(c *gin.Context) {
	var req bookReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data, err := services.CreateBook(
		req.Title, 
		req.Author, 
		req.Year, 
		req.Pages, 
		req.CategoryID)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, data)
}

func ListBooks(c *gin.Context) {
	data, err := services.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := services.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req bookReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	b := &models.Book{
		Title:      req.Title,
		Author:     req.Author,
		Year: 		req.Year,
		Pages: 		req.Pages,
		CategoryID: req.CategoryID,
	}

	if err := services.UpdateBook(id, b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteBook(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
