package models

import "time"

type Resource struct {
    ID        string    `json:"id" bson:"_id"`
    Title     string    `json:"title" bson:"title"`
    Type      string    `json:"type" bson:"type"`
    Audience  string    `json:"audience" bson:"audience"`
    FileName  string    `json:"fileName" bson:"filename"`
    FilePath  string    `json:"filePath" bson:"filepath"`
    FileSize  int64     `json:"fileSize" bson:"filesize"`
    DateAdded time.Time `json:"dateAdded" bson:"dateadded"`
}
