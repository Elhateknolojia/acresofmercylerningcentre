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
    db.InitResources()
    db.InitEvents()
    db.InitHighlights()


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

    // Resources routes
    r.POST("/api/resources", handlers.UploadResource)   // admin upload
    r.GET("/api/resources", handlers.ListResources)     // list all
    r.GET("/api/resources/:id/download", handlers.DownloadResource) // download
    r.DELETE("/api/resources/:id", handlers.DeleteResource) // delete



    // Events routes
r.GET("/api/events", handlers.ListEvents)
r.POST("/api/events", handlers.AddEvent)
r.DELETE("/api/events/:id", handlers.DeleteEvent)

// Highlights routes
r.GET("/api/highlights", handlers.ListHighlights)
r.POST("/api/highlights", handlers.SaveHighlights)
r.DELETE("/api/highlights/:id", handlers.DeleteHighlight)



    
    r.Run(":8080")
}
