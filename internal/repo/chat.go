package repo

import (
	"context"
	"strconv"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatRepository struct {
	collection *mongo.Collection
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	return &ChatRepository{
		collection: db.Collection("chat"),
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

func (r *ChatRepository) GetRoomMesages(roomID string, limit string) (*[]models.Message, error) {
	limitInt, err := strconv.Atoi(limit)
	roomObjID, err := primitive.ObjectIDFromHex(roomID)
	ctx := context.Background()
	if err != nil {
		return nil, err
	}
	filter := bson.M{"room_id": roomObjID}
	optns := options.Find().SetLimit(int64(limitInt)).SetSort(bson.M{"created_at": -1}) // Sort new first
	cursor, err := r.collection.Find(ctx, filter, optns)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []models.Message

	for cursor.Next(ctx) {
		var msg models.Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &messages, nil

}
