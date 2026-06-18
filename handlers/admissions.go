package handlers

import (
    "acresofmercy/models"
    "acresofmercy/utils"
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

func SubmitAdmission(c *gin.Context) {
    var req models.AdmissionRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    subject := "New Admission Request"
    body := "Name: " + req.ParentName + "\nEmail: " + req.Email + "\nPhone: " + req.Phone + "\nMessage: " + req.Message

    if err := utils.SendMail(
        "admissions@acresofmercylearningcentre.co.ke",
        subject,
        body,
        "MAIL_PASS_ADMISSIONS",
    ); err != nil {
        fmt.Println("SendMail error:", err) // log to Render
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"message": "Admission request submitted successfully"})
}
