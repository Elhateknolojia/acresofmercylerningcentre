package main

import (
	"acresofmercy/handlers"
	"acresofmercy/middleware"
	"os"
    "acresofmercy/db"
    "acresofmercy/services"
    "time"
	"github.com/gin-gonic/gin"
)

func main() {
    db.InitMongo(os.Getenv("MONGO_URI"))
        // Start background dispatcher every 30 minutes
    go func() {
        ticker := time.NewTicker(30 * time.Minute)
        defer ticker.Stop()
        for {
            <-ticker.C
            services.DispatchNewsletters()
        }
    }()

    r := gin.Default()

	 // Apply CORS middleware
    r.Use(middleware.CORSMiddleware())

    // Admissions route
    r.POST("/api/admissions", handlers.SubmitAdmission)

    // Contacts route
    r.POST("/api/contacts", handlers.SubmitContact)


    // new route
    r.POST("/api/subscribe", handlers.SubscribeHandler)

    r.Run(":8080")
}
