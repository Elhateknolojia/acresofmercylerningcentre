// handlers/newsletter.go
package handlers

import (
	"context"
	// "encoding/json"
	"net/http"
	"time"

	"acresofmercy/db"
	"acresofmercy/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func SubscribeHandler(c *gin.Context) {
    var req struct {
        Email string `json:"email"`
    }
    if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    coll := db.Client.Database("AOMLC").Collection("subscribers")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // prevent duplicates
    var existing models.Subscriber
    err := coll.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existing)
    if err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already subscribed"})
        return
    }

    sub := models.Subscriber{
        Email:     req.Email,
        CreatedAt: time.Now(),
    }

    _, err = coll.InsertOne(ctx, sub)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Subscribed successfully"})
}
// handlers/newsletter.go
func CreateNewsletter(c *gin.Context) {
    var req struct {
        Subject string `json:"subject"`
        Body    string `json:"body"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    coll := db.Client.Database("AOMLC").Collection("newsletters")

    _, err := coll.InsertOne(context.Background(), bson.M{
        "subject":   req.Subject,
        "body":      req.Body,
        "createdAt": time.Now(),
        "sent":      false,
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save newsletter"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Newsletter created"})
}
