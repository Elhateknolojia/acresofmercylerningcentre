// models/subscriber.go
package models

import "time"

type Subscriber struct {
    Email     string    `bson:"email"`
    CreatedAt time.Time `bson:"createdAt"`
}
