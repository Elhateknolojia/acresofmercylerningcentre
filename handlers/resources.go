package handlers

import (
    "acresofmercy/db"
    "acresofmercy/models"
    "net/http"
    "os"
     "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
    // "path/filepath"
	"time"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// UploadResource: admin uploads a file

func UploadResource(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }

    // Initialize Cloudinary
    cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))

    // Upload file
    f, _ := file.Open()
    uploadResult, err := cld.Upload.Upload(c, f, uploader.UploadParams{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to Cloudinary"})
        return
    }

    resource := models.Resource{
        ID:       uuid.New().String(),
        Title:    c.PostForm("title"),
        Type:     c.PostForm("type"),
        Audience: c.PostForm("audience"),
        FileName: file.Filename,
        FilePath: uploadResult.SecureURL, // ✅ Cloudinary URL
        FileSize: file.Size,
        DateAdded: time.Now(),
    }

    db.SaveResource(resource)
    c.JSON(http.StatusOK, resource)
}
// ListResources: return all resources
func ListResources(c *gin.Context) {
    resources := db.GetResources()
    c.JSON(http.StatusOK, resources)
}

// DownloadResource: stream file back
func DownloadResource(c *gin.Context) {
    id := c.Param("id")
    resource := db.GetResourceByID(id)
    if resource == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
        return
    }
    c.File(resource.FilePath)
}

// DeleteResource: remove from DB and disk
func DeleteResource(c *gin.Context) {
    id := c.Param("id")
    resource := db.GetResourceByID(id)
    if resource == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
        return
    }
    os.Remove(resource.FilePath)
    db.DeleteResource(id)
    c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
