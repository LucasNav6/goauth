package verification

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func Create(config *goauth_models.Configuration, identifier string, value string, expiresAtSeconds int64) (*goauth_entities.Verification, error) {
	id := utilities.GenerateUUID()
	expires := utilities.GetExpiryTimestamp(expiresAtSeconds)
	v, err := config.Entities.CreateVerification(*config.Context, goauth_entities.CreateVerificationParams{
		ID:         id,
		Identifier: identifier,
		Value:      value,
		Expiresat:  pgtype.Timestamptz{Time: expires, Valid: true},
		Createdat:  pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
		Updatedat:  pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &v, nil
}
