package controllers

import (
	"net/http"

	"example.com/sumit/database"
	"example.com/sumit/models"
	"github.com/gin-gonic/gin"
)

// Get all books

func FindBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Create New Book
func CreateBook(c *gin.Context) {
	// validate input
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// We first validate the request body by using the ShouldBindJSON method and pass the schema. If the data is invalid, it will return a 400 error to the client and tell them which fields are invalid. Otherwise, it will create a new book, save it to the database, and return the book.

// FindBook
func FindBook(c *gin.Context) {
	var book models.Book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook

func UpdateBook(c *gin.Context) {
	var book, input models.Book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
