package models

import (
	"time"
)

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updatedAt" binding:"required"`
}