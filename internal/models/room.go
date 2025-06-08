package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description" json:"description"`
	CreatorID   primitive.ObjectID   `bson:"creator_id" json:"creator_id"`
	Members     []primitive.ObjectID `bson:"members" json:"members"`
	LastMessage *Message             `bson:"-" json:"last_message,omitempty"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at" json:"updated_at"`
}

type CreateRoomRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=50"`
	Description string   `json:"description"`
	Members     []string `json:"members" validate:"required,min=1,dive,email"` // dive - перевірити кожен елемент масиву
}
