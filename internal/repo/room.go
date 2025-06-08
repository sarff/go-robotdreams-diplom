package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/models"
	log "github.com/sarff/iSlogger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomRepository struct {
	collection *mongo.Collection
}

func NewRoomRepository(db *mongo.Database) *RoomRepository {
	collection := db.Collection("room")

	indexName := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := collection.Indexes().CreateOne(context.Background(), indexName); err != nil {
		log.Warn("Failed to create index: ", "indexName", err)
	}

	return &RoomRepository{
		collection: collection,
	}
}

func (r *RoomRepository) CreateRoom(room *models.Room) error {
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

func (r *RoomRepository) FindByUserID(userId string) ([]*models.Room, error) {
	ctx := context.Background()
	userObjID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	cursor, err := r.collection.Find(ctx, bson.M{"members": userObjID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []*models.Room
	for cursor.Next(ctx) {
		var room models.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
