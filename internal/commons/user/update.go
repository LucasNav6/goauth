package user

import (
    "fmt"

    "github.com/LucasNav6/goauth/internal/utilities"
    goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
    goauth_models "github.com/LucasNav6/goauth/pkg/models"
    "github.com/jackc/pgx/v5/pgtype"
)

// Update updates mutable user fields.
func Update(config *goauth_models.Configuration, uuid string, name *string, email *string, image *string) error {
    if uuid == "" {
        return fmt.Errorf("uuid required")
    }

    // Fetch existing
    existing, err := config.Entities.GetUser(*config.Context, uuid)
    if err != nil {
        return err
    }
    if existing.ID == "" {
        return fmt.Errorf("user not found")
    }

    // Prepare fields
    var nameTxt pgtype.Text
    if name != nil {
        nameTxt = pgtype.Text{String: *name, Valid: true}
    } else {
        nameTxt = existing.Name
    }

    var emailTxt pgtype.Text
    if email != nil {
        if !utilities.IsValidEmail(*email) {
            return fmt.Errorf("invalid email")
        }
        emailTxt = pgtype.Text{String: *email, Valid: true}
    } else {
        emailTxt = existing.Email
    }

    var imageTxt pgtype.Text
    if image != nil {
        imageTxt = pgtype.Text{String: *image, Valid: true}
    } else {
        imageTxt = existing.Image
    }

    // Update
    err = config.Entities.UpdateUser(*config.Context, goauth_entities.UpdateUserParams{
        ID:            uuid,
        Name:          nameTxt,
        Email:         emailTxt,
        Emailverified: existing.Emailverified,
        Image:         imageTxt,
        Updatedat:     pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
    })
    return err
}
