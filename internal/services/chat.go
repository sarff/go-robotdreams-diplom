package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatService struct {
	cfg   *config.Config
	clnts *clients.Clients
	repo  *repo.ChatRepository
}

func NewChatService(cfg *config.Config, clnts *clients.Clients) *ChatService {
	return &ChatService{
		cfg:   cfg,
		clnts: clnts,
		repo:  repo.NewChatRepository(clnts.Mongo.DB),
	}
}

func (s *ChatService) SendMessage(userID string, req *models.MessageRequest) (*models.Message, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	roomObjID, err := primitive.ObjectIDFromHex(req.RoomID)
	if err != nil {
		return nil, err
	}

	message := &models.Message{
		RoomID:  roomObjID,
		UserID:  userObjID,
		Content: req.Content,
	}

	if err = s.repo.Create(message); err != nil {
		return nil, err
	}

	return message, nil
}
