package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	collection *mongo.Collection
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	return &ChatRepository{
		collection: db.Collection("room"),
	}
}

func (r *ChatRepository) CreateMessage(message *models.Message) error {
	message.CreatedAt = time.Now().UTC()
	message.UpdatedAt = time.Now().UTC()

	result, err := r.collection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}

	message.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *ChatRepository) CreateRoom(room *models.Room) error {
	room.CreatedAt = time.Now().UTC()
	room.UpdatedAt = time.Now().UTC()

	result, err := r.collection.InsertOne(context.Background(), room)
	if err != nil {
		var writeErr mongo.WriteException
		if errors.As(err, &writeErr) && writeErr.HasErrorCode(11000) {
			return fmt.Errorf("room with that name already exists")
		}
		return err
	}

	room.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
