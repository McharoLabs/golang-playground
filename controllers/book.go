package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mcharolabs/go-crud/database"
	"github.com/mcharolabs/go-crud/models"
	"github.com/mcharolabs/go-crud/utils/logger"
)

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		logger.Error("DeleteBook: Invalid UUID format - " + err.Error())
		return
	}

	var book models.Book

	if err := database.DB.Delete(&book, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book"})
		logger.Error("DeleteBook: Failed to delete book - " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		logger.Error("UpdateBook: Invalid UUID format - " + err.Error())
		return
	}

	var book models.Book

	if err := database.DB.First(&book, "id = ?", uuid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		logger.Error("UpdateBook: Book not found - " + err.Error())
		return
	}

	var input struct {
		Title  *string `json:"title"`
		Author *string `json:"author"`
		Year   *int    `json:"year"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		logger.Error("UpdateBook: Invalid input - " + err.Error())
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}
	if input.Author != nil {
		book.Author = *input.Author
	}
	if input.Year != nil {
		book.Year = *input.Year
	}

	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		logger.Error("UpdateBook: Failed to update book - " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "book": book})
	logger.Info("UpdateBook: Book updated successfully - " + book.ID.String())
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	uuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		logger.Error("BetBook: Invalid UUID format - " + err.Error())
		return
	}

	result := database.DB.First(&book, "id = ?", uuid)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		logger.Error("Book not found- " + result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {

	var books []models.Book

	result := database.DB.Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch books",
		})
		logger.Error("Failed to fetch books - " + result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		logger.Error("CreateBook: Invalid input - " + err.Error())
		return
	}

	result := database.DB.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		logger.Error("CreateBook: Failed to create book - " + result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book": book})
	logger.Info("CreateBook: Book created successfully - " + book.Title)
}
