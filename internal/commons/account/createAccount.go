package account

import (
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateAccountWithEmailAndPassword(config *goauth_models.Configuration, user goauth_models.UserUnauthenticated, userUuid string) (*goauth_models.Account, error) {
	uuid := utilities.GenerateUUID()
	hashedPassword, err := utilities.HashPassword(*user.Password)
	if err != nil {
		return nil, err
	}

	account, err := config.Entities.CreateAccount(*config.Context, goauth_entities.CreateAccountParams{
		ID:                    uuid,
		Userid:                userUuid,
		Accountid:             userUuid,
		Accesstoken:           pgtype.Text{Valid: false},
		Refreshtoken:          pgtype.Text{Valid: false},
		Accesstokenexpiresat:  pgtype.Timestamptz{Valid: false},
		Refreshtokenexpiresat: pgtype.Timestamptz{Valid: false},
		Scope:                 pgtype.Text{Valid: false},
		Idtoken:               pgtype.Text{Valid: false},
		Password:              pgtype.Text{String: hashedPassword, Valid: true},
		Providerid:            goauth_models.EMAIL_AND_PASSWORD,
		Createdat:             pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
		Updatedat:             pgtype.Timestamptz{Time: utilities.GetCurrentTimestamp(), Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return &goauth_models.Account{
		UUID:     account.ID,
		Provider: goauth_models.EMAIL_AND_PASSWORD,
		UserUUID: userUuid,
	}, nil
}

func GetAccountByUserIDAndProvider(config *goauth_models.Configuration, userID string, providerID string) (*goauth_models.Account, error) {
	accountEntity, err := config.Entities.GetAccountByProviderAndUserId(*config.Context, goauth_entities.GetAccountByProviderAndUserIdParams{
		Userid:     userID,
		Providerid: providerID,
	})
	if err != nil {
		return nil, err
	}
	if accountEntity.ID == "" {
		return nil, nil // Account not found
	}
	return &goauth_models.Account{
		UUID:     accountEntity.ID,
		Provider: accountEntity.Providerid,
		UserUUID: accountEntity.Userid,
		Password: accountEntity.Password.String,
	}, nil
}
