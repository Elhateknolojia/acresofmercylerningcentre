package models

type ContactRequest struct {
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required,email"`
    Phone   string `json:"phone"`
    Subject string `json:"subject" binding:"required"`
    Message string `json:"message" binding:"required"`
    Consent bool   `json:"consent" binding:"required"`
}
