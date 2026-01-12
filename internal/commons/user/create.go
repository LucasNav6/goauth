package user

import (
    "fmt"

    "github.com/LucasNav6/goauth/internal/utilities"
    goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
    goauth_models "github.com/LucasNav6/goauth/pkg/models"
    "github.com/jackc/pgx/v5/pgtype"
)

// Create creates a new user after validating inputs and checking duplicates.
func Create(config *goauth_models.Configuration, user goauth_models.UserUnauthenticated) (*goauth_models.UserAuthenticated, error) {
    if !utilities.IsValidEmail(user.Email) {
        return nil, fmt.Errorf("invalid email")
    }

    emailText := pgtype.Text{String: user.Email, Valid: true}
    existingUser, err := config.Entities.GetUserByEmail(*config.Context, emailText)
    if err == nil && existingUser.ID != "" {
        return &goauth_models.UserAuthenticated{Uuid: existingUser.ID, Email: existingUser.Email.String, Name: existingUser.Name.String}, nil
    }

    uuid := utilities.GenerateUUID()

    var image pgtype.Text
    if user.Image != nil {
        image = pgtype.Text{String: *user.Image, Valid: true}
    } else {
        image = pgtype.Text{Valid: false}
    }

    created, err := config.Entities.CreateUser(*config.Context, goauth_entities.CreateUserParams{
        ID:            uuid,
        Email:         pgtype.Text{String: user.Email, Valid: true},
        Name:          pgtype.Text{String: user.Name, Valid: true},
        Emailverified: false,
        Image:         image,
        Createdat:     pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
        Updatedat:     pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
    })
    if err != nil {
        return nil, err
    }

    return &goauth_models.UserAuthenticated{Uuid: created.ID, Email: created.Email.String, Name: created.Name.String}, nil
}
