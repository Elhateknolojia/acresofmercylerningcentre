package db

import (
    "acresofmercy/models"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var eventsCollection *mongo.Collection
var highlightsCollection *mongo.Collection

func InitEvents() {
    eventsCollection = Client.Database("acresdb").Collection("events")
}

func InitHighlights() {
    highlightsCollection = Client.Database("acresdb").Collection("highlights")
}

// --- Events ---
func GetEvents() []models.CalendarEvent {
    var events []models.CalendarEvent
    cursor, _ := eventsCollection.Find(context.TODO(), bson.M{})
    cursor.All(context.TODO(), &events)
    return events
}

func SaveEvent(event models.CalendarEvent) {
    eventsCollection.InsertOne(context.TODO(), event)
}

func DeleteEvent(id string) {
    eventsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
}

// --- Highlights ---
func GetHighlights() []models.Highlight {
    var highlights []models.Highlight
    cursor, _ := highlightsCollection.Find(context.TODO(), bson.M{})
    cursor.All(context.TODO(), &highlights)
    return highlights
}

func SaveHighlights(highlights []models.Highlight) {
    // Replace all highlights with new list
    highlightsCollection.Drop(context.TODO())
    for _, h := range highlights {
        highlightsCollection.InsertOne(context.TODO(), h)
    }
}

func GetHighlightByID(id string) *models.Highlight {
    var h models.Highlight
    err := highlightsCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&h)
    if err != nil {
        return nil
    }
    return &h
}

func DeleteHighlight(id string) {
    highlightsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
}
