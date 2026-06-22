package models

import "time"

type Resource struct {
    ID        string    `json:"id" bson:"_id"`
    Title     string    `json:"title"`
    Type      string    `json:"type"`
    Audience  string    `json:"audience"`
    FileName  string    `json:"fileName"`
    FilePath  string    `json:"filePath"`
    FileSize  int64     `json:"fileSize"`
    DateAdded time.Time `json:"dateAdded"`
}
