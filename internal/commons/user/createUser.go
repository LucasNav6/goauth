package user

import (
	"fmt"

	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetUserByEmail(config *goauth_models.Configuration, email string) (*goauth_models.UserAuthenticated, error) {
	emailText := pgtype.Text{String: email, Valid: true}
	existingUser, err := config.Entities.GetUserByEmail(*config.Context, emailText)
	if err != nil {
		return nil, err
	}
	if existingUser.ID == "" {
		return nil, fmt.Errorf("user not found")
	}
	return &goauth_models.UserAuthenticated{
		Uuid:  existingUser.ID,
		Email: existingUser.Email.String,
		Name:  existingUser.Name.String,
	}, nil
}

func CreateUser(config *goauth_models.Configuration, user goauth_models.UserUnauthenticated) (*goauth_models.UserAuthenticated, error) {
	// Check if user already exists
	emailText := pgtype.Text{String: user.Email, Valid: true}
	existingUser, err := config.Entities.GetUserByEmail(*config.Context, emailText)
	if err == nil && existingUser.ID != "" {
		err := fmt.Errorf("user with this email already exists")
		return nil, err
	}

	// If the user exists, return
	if err == nil && existingUser.ID != "" {
		return &goauth_models.UserAuthenticated{
			Uuid:  existingUser.ID,
			Email: user.Email,
			Name:  user.Name,
		}, nil
	}

	// Generate a new UUID for the user
	uuid := utilities.GenerateUUID()

	// Prepare image field
	var image pgtype.Text
	if user.Image != nil {
		image = pgtype.Text{String: *user.Image, Valid: true}
	} else {
		image = pgtype.Text{Valid: false}
	}

	// Create the user in the database
	config.Entities.CreateUser(*config.Context, goauth_entities.CreateUserParams{
		ID:            uuid,
		Email:         pgtype.Text{String: user.Email, Valid: true},
		Name:          pgtype.Text{String: user.Name, Valid: true},
		Emailverified: false,
		Image:         image,
		Createdat:     pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
		Updatedat:     pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})

	return &goauth_models.UserAuthenticated{
		Uuid:  uuid,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
