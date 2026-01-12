package session

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func Create(config *goauth_models.Configuration, userUUID string, ipAddress string, userAgent string) (*goauth_models.Session, error) {
	s, err := config.Entities.CreateSession(*config.Context, goauth_entities.CreateSessionParams{
		ID:        utilities.GenerateUUID(),
		Userid:    userUUID,
		Token:     utilities.GenerateToken(*config, userUUID),
		Expiresat: pgtype.Timestamptz{Time: utilities.GetExpiryTimestamp(config.SessionDurationInSeconds), Valid: true},
		Ipaddress: pgtype.Text{String: ipAddress, Valid: ipAddress != ""},
		Useragent: pgtype.Text{String: userAgent, Valid: userAgent != ""},
		Createdat: pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
		Updatedat: pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &goauth_models.Session{UUID: s.ID, UserUUID: s.Userid, Token: s.Token}, nil
}
