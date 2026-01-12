package account

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func UpdatePassword(config *goauth_models.Configuration, uuid string, newPassword string) error {
	if uuid == "" {
		return nil
	}
	hashed, err := utilities.HashPassword(newPassword)
	if err != nil {
		return err
	}
	// Fetch existing to preserve other fields
	existing, err := config.Entities.GetAccount(*config.Context, uuid)
	if err != nil {
		return err
	}
	if existing.ID == "" {
		return nil
	}

	return config.Entities.UpdateAccount(*config.Context, goauth_entities.UpdateAccountParams{
		ID:                    uuid,
		Userid:                existing.Userid,
		Accountid:             existing.Accountid,
		Providerid:            existing.Providerid,
		Accesstoken:           existing.Accesstoken,
		Refreshtoken:          existing.Refreshtoken,
		Accesstokenexpiresat:  existing.Accesstokenexpiresat,
		Refreshtokenexpiresat: existing.Refreshtokenexpiresat,
		Scope:                 existing.Scope,
		Idtoken:               existing.Idtoken,
		Password:              pgtype.Text{String: hashed, Valid: true},
		Updatedat:             pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})
}
