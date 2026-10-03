package user

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

var ErrInvalidUser = errors.New("name and phone are required")

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone,omitempty"`
	Guest     bool      `json:"guest"`
	CreatedAt time.Time `json:"createdAt"`
}

type Session struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Repository interface {
	Save(user User) (User, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Register(name string, phone string) (Session, error) {
	name = strings.TrimSpace(name)
	phone = strings.TrimSpace(phone)
	if name == "" || phone == "" {
		return Session{}, ErrInvalidUser
	}

	createdUser, err := service.repository.Save(User{
		ID:        "usr_" + randomHex(8),
		Name:      name,
		Phone:     phone,
		Guest:     false,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return Session{}, err
	}

	return Session{Token: "mock_" + randomHex(16), User: createdUser}, nil
}

func (service *Service) Guest() Session {
	return Session{
		Token: "guest_" + randomHex(16),
		User: User{
			ID:        "gst_" + randomHex(8),
			Name:      "Guest",
			Guest:     true,
			CreatedAt: time.Now().UTC(),
		},
	}
}

func randomHex(bytesCount int) string {
	bytes := make([]byte, bytesCount)
	if _, err := rand.Read(bytes); err != nil {
		return hex.EncodeToString([]byte(time.Now().UTC().Format(time.RFC3339Nano)))[:bytesCount]
	}
	return hex.EncodeToString(bytes)
}
