package account

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateWithPassword(config *goauth_models.Configuration, userUUID string, password string) (*goauth_models.Account, error) {
	uuid := utilities.GenerateUUID()
	hashed, err := utilities.HashPassword(password)
	if err != nil {
		return nil, err
	}

	acc, err := config.Entities.CreateAccount(*config.Context, goauth_entities.CreateAccountParams{
		ID:                    uuid,
		Userid:                userUUID,
		Accountid:             userUUID,
		Accesstoken:           pgtype.Text{Valid: false},
		Refreshtoken:          pgtype.Text{Valid: false},
		Accesstokenexpiresat:  pgtype.Timestamptz{Valid: false},
		Refreshtokenexpiresat: pgtype.Timestamptz{Valid: false},
		Scope:                 pgtype.Text{Valid: false},
		Idtoken:               pgtype.Text{Valid: false},
		Password:              pgtype.Text{String: hashed, Valid: true},
		Providerid:            goauth_models.EMAIL_AND_PASSWORD,
		Createdat:             pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
		Updatedat:             pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &goauth_models.Account{UUID: acc.ID, Provider: acc.Providerid, UserUUID: acc.Userid, Password: acc.Password.String}, nil
}
