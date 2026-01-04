package goauth

import (
	"context"
	"fmt"

	"github.com/LucasNav6/goauth/models"
	entites "github.com/LucasNav6/goauth/models/entities"
	"github.com/LucasNav6/goauth/utilities"
	"github.com/jackc/pgx/v5/pgtype"
)

func SignUpWithEmailAndPassword(cfg *models.Configuration, user entites.User, password string) (error, *entites.Account) {
	// Initialize queries to interact with the database
	ctx := context.Background()
	queries := entites.New(*cfg.EntitesDBTX)

	// Validate if the user already exists
	user, error := queries.GetUserByEmail(ctx, pgtype.Text{String: user.Email.String, Valid: true})
	if error == nil {
		return error, nil
	}

	// If the user does not exist, create a new user
	if user.ID == "" {
		newUser, error := queries.CreateUser(ctx, entites.CreateUserParams{
			Name:  pgtype.Text{String: user.Name.String, Valid: true},
			Email: pgtype.Text{String: user.Email.String, Valid: true},
			Image: pgtype.Text{String: user.Image.String, Valid: true},
		})
		if error != nil {
			return error, nil
		}

		user = newUser
	}

	// Create the authentication record for the user (EMAIL_AND_PASSWORD)
	account, error := queries.CreateAccount(ctx, entites.CreateAccountParams{
		ID:                    utilities.GenerateUUID(),
		Userid:                user.ID,
		Accountid:             utilities.GenerateUUID(),
		Providerid:            models.EMAIL_AND_PASSWORD,
		Password:              pgtype.Text{String: utilities.HashPassword(password), Valid: true},
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
		return error, nil
	}

	// Return the created account
	return nil, &account
}

func SignInWithEmailAndPassword(cfg *models.Configuration, email string, password string) (error, *entites.User, *entites.Session) {
	// Check if the user exists
	ctx := context.Background()
	queries := entites.New(*cfg.EntitesDBTX)

	user, error := queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if error != nil {
		return error, nil, nil
	}

	// Get the account associated with the user
	account, error := queries.GetAccountByProviderAndUserId(ctx, entites.GetAccountByProviderAndUserIdParams{
		Providerid: models.EMAIL_AND_PASSWORD,
		Userid:     user.ID,
	})
	if error != nil {
		return error, nil, nil
	}

	// Verify the password
	if !utilities.CheckPasswordHash(password, account.Password.String) {
		return fmt.Errorf("The credentials are invalid"), nil, nil
	}

	// Create the session for the user
	expiresAt := utilities.GetCurrentTimestampPlusSeconds(cfg.SessionExpirationInSeconds)
	session, error := queries.CreateSession(ctx, entites.CreateSessionParams{
		ID:        utilities.GenerateUUID(),
		Userid:    user.ID,
		Token:     utilities.GenerateUUID(),
		Expiresat: pgtype.Timestamptz{Time: utilities.TimestampToTime(expiresAt), Valid: true},
		Ipaddress: pgtype.Text{Valid: false},
		Useragent: pgtype.Text{Valid: false},
		Createdat: pgtype.Timestamptz{Valid: true},
		Updatedat: pgtype.Timestamptz{Valid: true},
	})
	if error != nil {
		return error, nil, nil
	}

	// Return the authenticated user
	return nil, &user, &session
}

func ResetPasswordWithEmailAndPassword(cfg *models.Configuration, email string, oldPassword string, newPassword string) error {
	// Check if the user exists
	ctx := context.Background()
	queries := entites.New(*cfg.EntitesDBTX)

	user, error := queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if error != nil {
		return error
	}

	// Get the account associated with the user
	account, error := queries.GetAccountByProviderAndUserId(ctx, entites.GetAccountByProviderAndUserIdParams{
		Providerid: models.EMAIL_AND_PASSWORD,
		Userid:     user.ID,
	})
	if error != nil {
		return error
	}

	// Verify the old password
	if !utilities.CheckPasswordHash(oldPassword, account.Password.String) {
		return fmt.Errorf("The credentials are invalid")
	}

	// Update the password to the new password
	error = queries.UpdateAccount(ctx, entites.UpdateAccountParams{
		ID:                    account.ID,
		Userid:                user.ID,
		Accountid:             account.Accountid,
		Providerid:            models.EMAIL_AND_PASSWORD,
		Password:              pgtype.Text{String: utilities.HashPassword(newPassword), Valid: true},
		Accesstoken:           pgtype.Text{Valid: false},
		Refreshtoken:          pgtype.Text{Valid: false},
		Accesstokenexpiresat:  pgtype.Timestamptz{Valid: false},
		Refreshtokenexpiresat: pgtype.Timestamptz{Valid: false},
		Scope:                 pgtype.Text{Valid: false},
		Idtoken:               pgtype.Text{Valid: false},
		Updatedat:             pgtype.Timestamptz{Valid: true},
	})
	if error != nil {
		return error
	}

	return nil
}
