package handlers

import (
    "acresofmercy/models"
    "acresofmercy/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

func SubmitAdmission(c *gin.Context) {
    var req models.AdmissionRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    subject := "New Admission Request"
    body := "Name: " + req.Name + "\nEmail: " + req.Email + "\nPhone: " + req.Phone + "\nMessage: " + req.Message

    if err := utils.SendMail("admissions@acresofmercylearningcentre.sc.ke", subject, body); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Admission request submitted successfully"})
}
