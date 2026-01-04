package goauth

import (
	"github.com/LucasNav6/goauth/models"
	entites "github.com/LucasNav6/goauth/models/entities"
	"github.com/LucasNav6/goauth/providers"
)

func CreateConfiguration() *models.Configuration {
	return &models.Configuration{}
}

func SignUpWithEmailAndPassword(cfg *models.Configuration, createUser models.ICreateUser) (error, *entites.Account) {
	return providers.SignUpWithEmailAndPassword(cfg, createUser)
}

func SignInWithEmailAndPassword(cfg *models.Configuration, email string, password string) (error, *entites.User, *entites.Session) {
	return providers.SignInWithEmailAndPassword(cfg, email, password)
}

func ResetPasswordWithEmailAndPassword(cfg *models.Configuration, email string, oldPassword string, newPassword string) error {
	return providers.ResetPasswordWithEmailAndPassword(cfg, email, oldPassword, newPassword)
}

func SignUpWithMagicLink(cfg *models.Configuration, email string) (error, *entites.Account) {
	return providers.SignUpWithMagicLink(cfg, email)
}

func SignInWithMagicLink(cfg *models.Configuration, email string, token string, expirationInSeconds int64) (error, *entites.Session) {
	return providers.SignInWithMagicLink(cfg, email, token, expirationInSeconds)
}

func ValidateMagicLinkSession(cfg *models.Configuration, token string) (error, *entites.Session) {
	return providers.ValidateMagicLinkSession(cfg, token)
}
