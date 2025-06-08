package services

import (
	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/repo"
	log "github.com/sarff/iSlogger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatService struct {
	cfg   *config.Config
	clnts *clients.Clients
	repo  *repo.Repos
}

func NewChatService(cfg *config.Config, clnts *clients.Clients, repos *repo.Repos) *ChatService {
	return &ChatService{
		cfg:   cfg,
		clnts: clnts,
		repo:  repos,
	}
}

func (s *ChatService) SendMessage(userID string, req *models.MessageRequest) (*models.Message, error) {
	//TODO: doesnt work
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

	if err = s.repo.ChatRepository.CreateMessage(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *ChatService) CreateRoom(userID string, m *models.CreateRoomRequest) (*models.Room, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	log.Debug("CreateRoom", "userObjID", userObjID)
	if err != nil {
		return nil, err
	}
	owner := false

	memberIDs := make([]primitive.ObjectID, 0) // why 0 ?
	// if we do len(m.members) with any non exist user
	//	"members": [
	//	"000000000000000000000000",
	//	"000000000000000000000000",
	//	"6844634bdcd820032ab52b09"
	//	],
	for _, memberID := range m.Members {
		log.Debug("CreateRoom", "memberID", memberID)
		usr, err := s.repo.UserRepository.FindByEmail(memberID)
		if err != nil {
			log.Error("CreateRoom", "FindByEmail", err)
			continue
		}

		if userObjID == usr.ID {
			owner = true
		}
		memberIDs = append(memberIDs, usr.ID)
	}

	if !owner {
		memberIDs = append(memberIDs, userObjID)
	}

	room := &models.Room{
		Name:        m.Name,
		Description: m.Description,
		CreatorID:   userObjID,
		Members:     memberIDs,
	}
	if err = s.repo.ChatRepository.CreateRoom(room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *ChatService) GetUserRooms(userID string) (interface{}, error) {
	// TODO: need implement
	userObjID, err := primitive.ObjectIDFromHex(userID)
	log.Debug("CreateRoom", "userObjID", userObjID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
