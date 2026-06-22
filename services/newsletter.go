// services/newsletter.go
package services

import (
	"acresofmercy/db"
	"acresofmercy/models"
	"acresofmercy/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsletterRequest struct {
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

func SendNewsletter(c *gin.Context) {
    var req NewsletterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := utils.RelayMail("https://acresofmercylearningcentre.sc.ke/sendnewsletter.php", req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Newsletter dispatched successfully"})
}

// services/newsletter.go
func DispatchNewsletters() {
    coll := db.Client.Database("AOMLC").Collection("newsletters")
    subs := db.Client.Database("AOMLC").Collection("subscribers")

    ctx := context.Background()

    cursor, _ := coll.Find(ctx, bson.M{"sent": false})
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var nl struct {
            ID      primitive.ObjectID `bson:"_id"`
            Subject string             `bson:"subject"`
            Body    string             `bson:"body"`
        }
        cursor.Decode(&nl)

        // Fetch subscribers
        subCursor, _ := subs.Find(ctx, bson.M{})
        defer subCursor.Close(ctx)

        for subCursor.Next(ctx) {
            var sub models.Subscriber
            subCursor.Decode(&sub)

            // Relay to PHP mailer
            payload := map[string]string{
                "subject": nl.Subject,
                "body":    nl.Body,
                "to":      sub.Email,
            }
            utils.RelayMail("https://acresofmercylearningcentre.sc.ke/sendnewsletter.php", payload)
        }

        // Mark newsletter as sent
        coll.UpdateOne(ctx, bson.M{"_id": nl.ID}, bson.M{"$set": bson.M{"sent": true}})
    }
}
