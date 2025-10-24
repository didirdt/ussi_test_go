package controllers

import (
	"net/http"
	"strconv"
	"ussi_test/database"
	"ussi_test/models"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	var notes []models.Note
	query := database.DB.Preload("User")

	if role == models.RoleUser {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req models.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := models.Note{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := database.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func GetNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	var note models.Note
	query := database.DB.Preload("User")

	if role == models.RoleUser {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	var note models.Note
	query := database.DB

	if role == models.RoleUser {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var req models.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Title != "" {
		note.Title = req.Title
	}
	if req.Content != "" {
		note.Content = req.Content
	}

	if err := database.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	var note models.Note
	query := database.DB

	if role == models.RoleUser {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if err := database.DB.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
