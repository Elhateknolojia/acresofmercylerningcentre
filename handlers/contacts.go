package handlers

import (
    "acresofmercy/models"
    "acresofmercy/utils"
    "net/http"
	"fmt"

    "github.com/gin-gonic/gin"
)

func SubmitContact(c *gin.Context) {
    var req models.ContactRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    subject := "New Contact Request - " + req.Subject
    body := "Name: " + req.Name +
        "\nEmail: " + req.Email +
        "\nPhone: " + req.Phone +
        "\nSubject: " + req.Subject +
        "\nMessage: " + req.Message +
        "\nConsent: " + fmt.Sprintf("%t", req.Consent)

    if err := utils.SendMail("contact@acresofmercylearningcentre.sc.ke", subject, body); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send contact request"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Contact request submitted successfully"})
}
