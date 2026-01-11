package session

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateSession(config *goauth_models.Configuration, userUUID string, ipAddress string, userAgent string) (*goauth_models.Session, error) {
	// Create the session in the database
	session, err := config.Entities.CreateSession(*config.Context, goauth_entities.CreateSessionParams{
		ID:        utilities.GenerateUUID(),
		Userid:    userUUID,
		Token:     utilities.GenerateToken(*config, userUUID),
		Expiresat: pgtype.Timestamptz{Time: utilities.GetExpiryTimestamp(config.SessionDurationInSeconds), Valid: true}, // Convert time.Time to pgtype.Timestamptz
		Ipaddress: pgtype.Text{String: ipAddress, Valid: true},                                                          // Assuming a utility function to get IP address
		Useragent: pgtype.Text{String: userAgent, Valid: true},                                                          // Assuming a utility function to get user agent
		Createdat: pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},                               // Assuming a utility function to get current timestamp
		Updatedat: pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},                               // Assuming a utility function to get current timestamp
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// Return the created session
	return &goauth_models.Session{
		UUID:     session.ID,
		UserUUID: userUUID,
	}, nil
}
