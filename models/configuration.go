package models

import (
	"context"

	entites "github.com/LucasNav6/goauth/models/entities"
)

type Configuration struct {
	Entites                    *entites.Queries
	Context                    *context.Context
	SessionExpirationInSeconds int64
	SendEmailCallback          func(toEmail string, subject string, body string) error
}
