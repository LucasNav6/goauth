package user

import (
	"fmt"

	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetAll returns all users.
func GetAll(config *goauth_models.Configuration) ([]goauth_models.UserAuthenticated, error) {
	ents, err := config.Entities.ListUsers(*config.Context)
	if err != nil {
		return nil, err
	}
	var out []goauth_models.UserAuthenticated
	for _, u := range ents {
		out = append(out, goauth_models.UserAuthenticated{Uuid: u.ID, Email: u.Email.String, Name: u.Name.String})
	}
	return out, nil
}

// GetByUUID returns a user by UUID.
func GetByUUID(config *goauth_models.Configuration, uuid string) (*goauth_models.UserAuthenticated, error) {
	if uuid == "" {
		return nil, fmt.Errorf("uuid required")
	}
	u, err := config.Entities.GetUser(*config.Context, uuid)
	if err != nil {
		return nil, err
	}
	if u.ID == "" {
		return nil, fmt.Errorf("user not found")
	}
	return &goauth_models.UserAuthenticated{Uuid: u.ID, Email: u.Email.String, Name: u.Name.String}, nil
}

// GetByEmail returns a user by email.
func GetByEmail(config *goauth_models.Configuration, email string) (*goauth_models.UserAuthenticated, error) {
	if !utilities.IsValidEmail(email) {
		return nil, fmt.Errorf("invalid email")
	}
	emailText := pgtype.Text{String: email, Valid: true}
	u, err := config.Entities.GetUserByEmail(*config.Context, emailText)
	if err != nil {
		return nil, err
	}
	if u.ID == "" {
		return nil, fmt.Errorf("user not found")
	}
	return &goauth_models.UserAuthenticated{Uuid: u.ID, Email: u.Email.String, Name: u.Name.String}, nil
}
