package handlers

import (
    "acresofmercy/db"
    "acresofmercy/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func ListHighlights(c *gin.Context) {
    highlights := db.GetHighlights()
    c.JSON(http.StatusOK, highlights)
}

func SaveHighlights(c *gin.Context) {
    var highlights []models.Highlight
    if err := c.ShouldBindJSON(&highlights); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.SaveHighlights(highlights)
    c.JSON(http.StatusOK, highlights)
}

func DeleteHighlight(c *gin.Context) {
    id := c.Param("id")
    highlight := db.GetHighlightByID(id)
    if highlight == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Highlight not found"})
        return
    }
    db.DeleteHighlight(id)
    c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
