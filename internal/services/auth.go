package services

import (
	"errors"
	"time"

	"github.com/sarff/go-robotdreams-diplom/internal/clients"
	"github.com/sarff/go-robotdreams-diplom/internal/config"
	"github.com/sarff/go-robotdreams-diplom/internal/models"
	"github.com/sarff/go-robotdreams-diplom/internal/repo"
	"github.com/sarff/go-robotdreams-diplom/internal/utils"
	log "github.com/sarff/iSlogger"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg   *config.Config
	clnts *clients.Clients
	repo  *repo.Repos
}

func NewAuthService(cfg *config.Config, clnts *clients.Clients, repos *repo.Repos) *AuthService {
	return &AuthService{
		cfg:   cfg,
		clnts: clnts,
		repo:  repos,
	}
}

func (as *AuthService) Register(req *models.RegisterRequest) error {
	existingUser, _ := as.repo.UserRepository.FindByEmail(req.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		IsOnline: true,
		LastSeen: time.Now(),
	}

	if err = as.repo.UserRepository.Create(user); err != nil {
		return err
	}

	return nil
}

func (as *AuthService) Login(req *models.LoginRequest) (*models.User, string, error) {
	user, err := as.repo.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Update online status
	err = as.repo.UserRepository.UpdateOnlineStatus(user.ID.Hex(), true)
	if err != nil {
		log.Error("failed to update online status", "error", err)
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID.Hex(), as.cfg.JWT.Secret)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (as *AuthService) Logout() error {
	// TODO: need implement
	return nil
}

func (as *AuthService) LogoutAll() error {
	// TODO: need implement
	return nil
}

func (as *AuthService) FindByID(userID string) (*models.User, error) {
	return as.repo.UserRepository.FindByID(userID)
}
