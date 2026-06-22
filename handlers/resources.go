package handlers

import (
    "acresofmercy/db"
    "acresofmercy/models"
    "net/http"
    "os"
    "path/filepath"
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

    id := uuid.New().String()
    savePath := filepath.Join("uploads", id+"_"+file.Filename)
    if err := c.SaveUploadedFile(file, savePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    resource := models.Resource{
        ID:       id,
        Title:    c.PostForm("title"),
        Type:     c.PostForm("type"),
        Audience: c.PostForm("audience"),
        FileName: file.Filename,
        FilePath: savePath,
        FileSize: file.Size,
        DateAdded: time.Now(),
    }

    db.SaveResource(resource) // persist in Mongo
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
