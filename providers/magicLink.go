package providers

import (
	"fmt"
	"strings"

	"github.com/LucasNav6/goauth/models"
	entites "github.com/LucasNav6/goauth/models/entities"
	"github.com/LucasNav6/goauth/utilities"
	"github.com/jackc/pgx/v5/pgtype"
)

func SignUpWithMagicLink(cfg *models.Configuration, email string) (*entites.Account, error) {
	// Validate if the user already exists
	user, error := cfg.Entites.GetUserByEmail(*cfg.Context, pgtype.Text{String: email, Valid: true})
	if error == nil {
		return nil, error
	}

	// If the user does not exist, create a new user
	if user.ID == "" {
		newUser, error := cfg.Entites.CreateUser(*cfg.Context, entites.CreateUserParams{
			Name:      pgtype.Text{String: strings.Split(email, "@")[0], Valid: true},
			Email:     pgtype.Text{String: email, Valid: true},
			Image:     pgtype.Text{String: "", Valid: false},
			Createdat: pgtype.Timestamptz{Valid: true},
			Updatedat: pgtype.Timestamptz{Valid: true},
		})
		if error != nil {
			return nil, error
		}

		user = newUser
	}

	// Create the authentication record for the user (EMAIL_AND_PASSWORD)
	account, error := cfg.Entites.CreateAccount(*cfg.Context, entites.CreateAccountParams{
		ID:                    utilities.GenerateUUID(),
		Userid:                user.ID,
		Accountid:             utilities.GenerateUUID(),
		Providerid:            models.MAGIC_LINK,
		Password:              pgtype.Text{Valid: false},
		Accesstoken:           pgtype.Text{Valid: false},
		Refreshtoken:          pgtype.Text{Valid: false},
		Accesstokenexpiresat:  pgtype.Timestamptz{Valid: false},
		Refreshtokenexpiresat: pgtype.Timestamptz{Valid: false},
		Scope:                 pgtype.Text{Valid: false},
		Idtoken:               pgtype.Text{Valid: false},
		Createdat:             pgtype.Timestamptz{Valid: true},
		Updatedat:             pgtype.Timestamptz{Valid: true},
	})
	if error != nil {
		return nil, error
	}

	return &account, nil
}

func SignInWithMagicLink(cfg *models.Configuration, email string, token string, expirationInSeconds int64) (*entites.Session, error) {
	// Check if the user exists
	user, error := cfg.Entites.GetUserByEmail(*cfg.Context, pgtype.Text{String: email, Valid: true})
	if error != nil {
		return nil, error
	}

	// Get the account associated with the user
	account, error := cfg.Entites.GetAccountByProviderAndUserId(*cfg.Context, entites.GetAccountByProviderAndUserIdParams{
		Providerid: models.MAGIC_LINK,
		Userid:     user.ID,
	})
	if error != nil {
		return nil, error
	}

	if account.ID == "" {
		return nil, fmt.Errorf("No MAGIC_LINK account found for this user")
	}

	// Create a new session for the user (temporal)
	expiresAt := utilities.GetCurrentTimestampPlusSeconds(expirationInSeconds)
	session, error := cfg.Entites.CreateSession(*cfg.Context, entites.CreateSessionParams{
		ID:        utilities.GenerateUUID(),
		Userid:    user.ID,
		Token:     token,
		Expiresat: pgtype.Timestamptz{Time: utilities.TimestampToTime(expiresAt), Valid: true},
		Ipaddress: pgtype.Text{Valid: false},
		Useragent: pgtype.Text{Valid: false},
		Createdat: pgtype.Timestamptz{Valid: true},
		Updatedat: pgtype.Timestamptz{Valid: true},
	})
	if error != nil {
		return nil, error
	}

	return &session, nil
}

func ValidateMagicLinkSession(cfg *models.Configuration, token string) (*entites.Session, error) {
	session, error := cfg.Entites.GetSessionByToken(*cfg.Context, token)
	if error != nil {
		return nil, error
	}

	// Update the session to be a permanent one
	expiresAt := utilities.GetCurrentTimestampPlusSeconds(cfg.SessionExpirationInSeconds)
	error = cfg.Entites.UpdateSession(*cfg.Context, entites.UpdateSessionParams{
		ID:        session.ID,
		Userid:    session.Userid,
		Token:     session.Token,
		Expiresat: pgtype.Timestamptz{Time: utilities.TimestampToTime(expiresAt), Valid: true},
		Ipaddress: session.Ipaddress,
		Useragent: session.Useragent,
		Updatedat: pgtype.Timestamptz{Valid: true},
	})

	// Return the session
	return &session, nil
}
