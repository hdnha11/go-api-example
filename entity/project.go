package entity

import (
	"time"
)

// Project entity
type Project struct {
	ID          ID        `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
