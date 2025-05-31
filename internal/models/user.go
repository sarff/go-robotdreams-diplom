package models

import "time"

type User struct {
	Username  string    `bson:"username" json:"username" validate:"required,min=3,max=20"`
	Email     string    `bson:"email" json:"email" validate:"required,email"`
	Password  string    `bson:"password" json:"-"`
	IsOnline  bool      `bson:"is_online" json:"is_online"`
	LastSeen  time.Time `bson:"last_seen" json:"last_seen"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
