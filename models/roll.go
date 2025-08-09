package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Struct da jogada do dado
type Roll struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Player    string             `json:"player" bson:"player"`
	Result    int                `json:"result" bson:"result"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

var collection *mongo.Collection
