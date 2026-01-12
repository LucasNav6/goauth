package verification

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func Update(config *goauth_models.Configuration, id string, identifier string, value string, expiresAtSeconds int64) error {
	if id == "" {
		return nil
	}
	expires := utilities.GetExpiryTimestamp(expiresAtSeconds)
	return config.Entities.UpdateVerification(*config.Context, goauth_entities.UpdateVerificationParams{
		ID:         id,
		Identifier: identifier,
		Value:      value,
		Expiresat:  pgtype.Timestamptz{Time: expires, Valid: true},
		Updatedat:  pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})
}
