package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repos struct {
	ChatRepository *ChatRepository
	UserRepository *UserRepository
	RoomRepository *RoomRepository
}

func NewRepos(db *mongo.Database) *Repos {
	return &Repos{
		ChatRepository: NewChatRepository(db),
		UserRepository: NewUserRepository(db),
		RoomRepository: NewRoomRepository(db),
	}
}
