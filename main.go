package main

import (
	"acresofmercy/middleware"
    "acresofmercy/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

	 // Apply CORS middleware
    r.Use(middleware.CORSMiddleware())

    // Admissions route
    r.POST("/api/admissions", handlers.SubmitAdmission)

    // Contacts route
    r.POST("/api/contacts", handlers.SubmitContact)

    r.Run(":8080")
}
