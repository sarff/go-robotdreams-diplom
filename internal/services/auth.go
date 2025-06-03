package services

import (
	"errors"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/repo"
	"github.com/sarff/go-robotdreams-diplom/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg   *config.Config
	clnts *clients.Clients
	repo  *repo.UserRepository
}

func NewAuthService(cfg *config.Config, clnts *clients.Clients) *AuthService {
	return &AuthService{
		cfg:   cfg,
		clnts: clnts,
	}
}

func (as *AuthService) Register(req *models.RegisterRequest) error {
	existingUser, _ := as.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		IsOnline: true,
		LastSeen: time.Now(),
	}

	if err = as.repo.Create(user); err != nil {
		return err
	}

	return nil
}

func (as *AuthService) Login(req *models.LoginRequest) (*models.User, string, error) {
	user, err := as.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Update online status TODO:
	//as.repo.UpdateOnlineStatus(user.ID.Hex(), true)

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID.Hex(), as.cfg.JWT.Secret)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (as *AuthService) Logout() error {
	return nil
}
