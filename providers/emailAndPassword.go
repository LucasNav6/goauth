package providers

import (
	"fmt"

	"github.com/LucasNav6/goauth/models"
	entites "github.com/LucasNav6/goauth/models/entities"
	"github.com/LucasNav6/goauth/utilities"
	"github.com/jackc/pgx/v5/pgtype"
)

func SignUpWithEmailAndPassword(cfg *models.Configuration, createUser models.ICreateUser) (*entites.Account, error) {
	// Validate if the user already exists
	user, error := cfg.Entites.GetUserByEmail(*cfg.Context, pgtype.Text{String: createUser.Email.String, Valid: true})
	if error == nil {
		return nil, fmt.Errorf("User with this email already exists")
	}

	// If the user does not exist, create a new user
	if user.ID == "" {
		newUser, error := cfg.Entites.CreateUser(*cfg.Context, entites.CreateUserParams{
			Name:  pgtype.Text{String: user.Name.String, Valid: true},
			Email: pgtype.Text{String: user.Email.String, Valid: true},
			Image: pgtype.Text{String: user.Image.String, Valid: true},
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
		Providerid:            models.EMAIL_AND_PASSWORD,
		Password:              pgtype.Text{String: utilities.HashPassword(createUser.Password), Valid: true},
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

	// Return the created account
	return &account, nil
}

func SignInWithEmailAndPassword(cfg *models.Configuration, email string, password string) (*entites.User, *entites.Session, error) {
	// Check if the user exists
	user, error := cfg.Entites.GetUserByEmail(*cfg.Context, pgtype.Text{String: email, Valid: true})
	if error != nil {
		return nil, nil, error
	}

	// Get the account associated with the user
	account, error := cfg.Entites.GetAccountByProviderAndUserId(*cfg.Context, entites.GetAccountByProviderAndUserIdParams{
		Providerid: models.EMAIL_AND_PASSWORD,
		Userid:     user.ID,
	})
	if error != nil {
		return nil, nil, error
	}

	// Verify the password
	if !utilities.CheckPasswordHash(password, account.Password.String) {
		return nil, nil, fmt.Errorf("The credentials are invalid")
	}

	// Create the session for the user
	expiresAt := utilities.GetCurrentTimestampPlusSeconds(cfg.SessionExpirationInSeconds)
	session, error := cfg.Entites.CreateSession(*cfg.Context, entites.CreateSessionParams{
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
		return nil, nil, error
	}

	// Return the authenticated user
	return &user, &session, nil
}

func ResetPasswordWithEmailAndPassword(cfg *models.Configuration, email string, oldPassword string, newPassword string) error {
	user, error := cfg.Entites.GetUserByEmail(*cfg.Context, pgtype.Text{String: email, Valid: true})
	if error != nil {
		return error
	}

	// Get the account associated with the user
	account, error := cfg.Entites.GetAccountByProviderAndUserId(*cfg.Context, entites.GetAccountByProviderAndUserIdParams{
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
	error = cfg.Entites.UpdateAccount(*cfg.Context, entites.UpdateAccountParams{
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
