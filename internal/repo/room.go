package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository struct {
	collection *mongo.Collection
}

func NewRoomRepository(db *mongo.Database) *RoomRepository {
	return &RoomRepository{
		collection: db.Collection("room"),
	}
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

func (r *RoomRepository) FindRoomByID(roomID string) (*models.Room, error) {
	roomObjectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, err
	}

	var room models.Room
	err = r.collection.FindOne(context.Background(), bson.M{"_id": roomObjectID}).Decode(&room)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}

	return &room, nil
}

func (r *RoomRepository) FindRoomByName(roomName string) (*models.Room, error) {
	var room models.Room
	err := r.collection.FindOne(context.Background(), bson.M{"name": roomName}).Decode(&room)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}

	return &room, nil
}
