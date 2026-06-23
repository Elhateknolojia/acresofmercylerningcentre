package handlers

import (
    "acresofmercy/db"
    "acresofmercy/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func ListEvents(c *gin.Context) {
    events := db.GetEvents()
    c.JSON(http.StatusOK, events)
}

func AddEvent(c *gin.Context) {
    var event models.CalendarEvent
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.SaveEvent(event)
    c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
    id := c.Param("id")
    db.DeleteEvent(id)
    c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
