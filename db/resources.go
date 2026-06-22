package db

import (
    "acresofmercy/models"
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
)

var resourceCollection *mongo.Collection

// InitMongo already sets up client in your db package.
func InitResources() {
    if Client == nil {
        log.Fatal("[DB] Mongo client not initialized. Call InitMongo first.")
    }
    db := Client.Database("acresofmercy")
    resourceCollection = db.Collection("resources")
    log.Println("[DB] Resources collection initialized")
}


// SaveResource inserts a new resource
func SaveResource(resource models.Resource) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    log.Printf("[DB] Saving resource: %+v\n", resource)

    _, err := resourceCollection.InsertOne(ctx, resource)
    if err != nil {
        log.Printf("[DB] Error inserting resource: %v\n", err)
    } else {
        log.Println("[DB] Resource saved successfully")
    }
}

// GetResources returns all resources
func GetResources() []models.Resource {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    log.Println("[DB] Fetching all resources")

    cursor, err := resourceCollection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("[DB] Error fetching resources: %v\n", err)
        return nil
    }
    defer cursor.Close(ctx)

    var resources []models.Resource
    if err := cursor.All(ctx, &resources); err != nil {
        log.Printf("[DB] Error decoding resources: %v\n", err)
        return nil
    }

    log.Printf("[DB] Found %d resources\n", len(resources))
    return resources
}

// GetResourceByID returns a single resource
func GetResourceByID(id string) *models.Resource {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    log.Printf("[DB] Fetching resource by ID: %s\n", id)

    var resource models.Resource
    err := resourceCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&resource)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            log.Printf("[DB] No resource found with ID: %s\n", id)
            return nil
        }
        log.Printf("[DB] Error fetching resource: %v\n", err)
        return nil
    }

    log.Printf("[DB] Resource found: %+v\n", resource)
    return &resource
}

// DeleteResource removes a resource
func DeleteResource(id string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    log.Printf("[DB] Deleting resource with ID: %s\n", id)

    res, err := resourceCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        log.Printf("[DB] Error deleting resource: %v\n", err)
        return
    }

    if res.DeletedCount == 0 {
        log.Printf("[DB] No resource deleted, ID not found: %s\n", id)
    } else {
        log.Printf("[DB] Resource deleted successfully, ID: %s\n", id)
    }
}
