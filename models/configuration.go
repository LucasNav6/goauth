package models

import entites "github.com/LucasNav6/goauth/models/entities"

type Configuration struct {
	EntitesDBTX                *entites.DBTX
	SessionExpirationInSeconds int64
	SendEmailCallback          func(toEmail string, subject string, body string) error
}
