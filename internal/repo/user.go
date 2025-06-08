package repo

import (
	"context"
	"errors"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/models"
	log "github.com/sarff/iSlogger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("room")

	indexName := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := collection.Indexes().CreateOne(context.Background(), indexName); err != nil {
		log.Warn("Failed to create index: ", "indexName", err)
	}

	return &UserRepository{
		collection: collection,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	// user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateOnlineStatus(userID string, isOnline bool) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	upd := bson.M{
		"$set": bson.M{
			"is_online":  isOnline,
			"updated_at": time.Now(),
			"last_seen":  time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, upd)

	return err
}
