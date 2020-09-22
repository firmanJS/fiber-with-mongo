package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"purpose_message,omitempty" binding:"required"`
	Email     string             `json:"email,omitempty" bson:"message,omitempty" binding:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" binding:"required"`
}
