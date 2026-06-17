package models

type AdmissionRequest struct {
    Name    string `json:"name" binding:"required"`
    Email   string `json:"email" binding:"required,email"`
    Phone   string `json:"phone" binding:"required"`
    Message string `json:"message" binding:"required"`
}
