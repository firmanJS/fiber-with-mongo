package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty" binding:"required"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" binding:"required"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" binding:"required"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" binding:"required"`
}
