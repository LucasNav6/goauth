package goauth

import (
	"github.com/LucasNav6/goauth/models"
	entites "github.com/LucasNav6/goauth/models/entities"
	"github.com/LucasNav6/goauth/providers"
)

func CreateConfiguration() *models.Configuration {
	return &models.Configuration{}
}

func SignUpWithEmailAndPassword(cfg *models.Configuration, createUser models.ICreateUser) (*entites.Account, error) {
	return providers.SignUpWithEmailAndPassword(cfg, createUser)
}

func SignInWithEmailAndPassword(cfg *models.Configuration, email string, password string) (*entites.User, *entites.Session, error) {
	return providers.SignInWithEmailAndPassword(cfg, email, password)
}

func ResetPasswordWithEmailAndPassword(cfg *models.Configuration, email string, oldPassword string, newPassword string) error {
	return providers.ResetPasswordWithEmailAndPassword(cfg, email, oldPassword, newPassword)
}

func SignUpWithMagicLink(cfg *models.Configuration, email string) (*entites.Account, error) {
	return providers.SignUpWithMagicLink(cfg, email)
}

func SignInWithMagicLink(cfg *models.Configuration, email string, token string, expirationInSeconds int64) (*entites.Session, error) {
	return providers.SignInWithMagicLink(cfg, email, token, expirationInSeconds)
}

func ValidateMagicLinkSession(cfg *models.Configuration, token string) (*entites.Session, error) {
	return providers.ValidateMagicLinkSession(cfg, token)
}
