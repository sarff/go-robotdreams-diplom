package services

import (
	"errors"

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
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	roomObjID, err := primitive.ObjectIDFromHex(req.RoomID)
	if err != nil {
		return nil, err
	}

	room, err := s.repo.RoomRepository.FindRoomByID(req.RoomID)
	if err != nil {
		return nil, err
	}

	isMember := false
	for _, memberID := range room.Members {
		if memberID == userObjID {
			isMember = true
			break
		}
	}
	if !isMember {
		return nil, errors.New("user is not a member of this room")
	}

	message := &models.Message{
		RoomID:  roomObjID,
		UserID:  userObjID,
		Content: req.Content,
	}

	if err = s.repo.ChatRepository.CreateMessage(message); err != nil {
		return nil, err
	}

	user, _ := s.repo.UserRepository.FindByID(userID)
	message.User = user

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
	for _, memberEmail := range m.Members {
		log.Debug("CreateRoom", "memberEmail", memberEmail)
		usr, err := s.repo.UserRepository.FindByEmail(memberEmail)
		if err != nil {
			log.Error("CreateRoom", "FindByEmail", err)
			continue
		}

		if userObjID == usr.ID {
			owner = true
		}
		memberIDs = append(memberIDs, usr.ID)
		log.Debug("CreateRoom", "memberIDs", memberIDs)
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
	if err = s.repo.RoomRepository.CreateRoom(room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *ChatService) GetUserRooms(userID string) ([]*models.Room, error) {
	return s.repo.RoomRepository.FindByUserID(userID)
}

func (s *ChatService) FindRoomByID(roomID string) (*models.Room, error) {
	return s.repo.RoomRepository.FindRoomByID(roomID)
}
func (s *ChatService) FindRoomByName(roomName string) (*models.Room, error) {
	return s.repo.RoomRepository.FindRoomByName(roomName)
}

func (s *ChatService) GetRoomMesages(id string, limit string) (*[]models.Message, error) {
	return s.repo.ChatRepository.GetRoomMesages(id, limit)
}
