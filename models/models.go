package models

type CalendarEvent struct {
    ID      string `json:"id" bson:"_id"`
    Title   string `json:"title" bson:"title"`
    DateStr string `json:"dateStr" bson:"dateStr"`
    Type    string `json:"type" bson:"type"`
}

type Highlight struct {
    ID          string `json:"id" bson:"_id"`
    Title       string `json:"title" bson:"title"`
    DateStr     string `json:"dateStr" bson:"dateStr"`
    Description string `json:"description" bson:"description"`
    ImageUrl    string `json:"imageUrl" bson:"imageUrl"`
    Category    string `json:"category" bson:"category"`
    Icon        string `json:"icon" bson:"icon"`
}
