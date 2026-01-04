package providers

import (
	"context"
	"fmt"
	"strings"

	"github.com/LucasNav6/goauth/models"
	entites "github.com/LucasNav6/goauth/models/entities"
	"github.com/LucasNav6/goauth/utilities"
	"github.com/jackc/pgx/v5/pgtype"
)

func SignUpWithMagicLink(cfg *models.Configuration, email string) (*entites.Account, error) {
	// Initialize queries to interact with the database
	ctx := context.Background()
	queries := entites.New(*cfg.EntitesDBTX)

	// Validate if the user already exists
	user, error := queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if error == nil {
		return nil, error
	}

	// If the user does not exist, create a new user
	if user.ID == "" {
		newUser, error := queries.CreateUser(ctx, entites.CreateUserParams{
			Name:  pgtype.Text{String: strings.Split(email, "@")[0], Valid: true},
			Email: pgtype.Text{String: email, Valid: true},
			Image: pgtype.Text{String: "", Valid: false},
		})
		if error != nil {
			return nil, error
		}

		user = newUser
	}

	// Create the authentication record for the user (EMAIL_AND_PASSWORD)
	account, error := queries.CreateAccount(ctx, entites.CreateAccountParams{
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
	ctx := context.Background()
	queries := entites.New(*cfg.EntitesDBTX)

	user, error := queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if error != nil {
		return nil, error
	}

	// Get the account associated with the user
	account, error := queries.GetAccountByProviderAndUserId(ctx, entites.GetAccountByProviderAndUserIdParams{
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
	session, error := queries.CreateSession(ctx, entites.CreateSessionParams{
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
	// Check if the session exists
	ctx := context.Background()
	queries := entites.New(*cfg.EntitesDBTX)

	session, error := queries.GetSessionByToken(ctx, token)
	if error != nil {
		return nil, error
	}

	// Update the session to be a permanent one
	expiresAt := utilities.GetCurrentTimestampPlusSeconds(cfg.SessionExpirationInSeconds)
	error = queries.UpdateSession(ctx, entites.UpdateSessionParams{
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
