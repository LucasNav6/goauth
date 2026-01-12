package account

import (
	"fmt"

	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func GetByUserAndProvider(config *goauth_models.Configuration, userID string, providerID string) (*goauth_models.Account, error) {
	if userID == "" || providerID == "" {
		return nil, fmt.Errorf("userID and providerID required")
	}
	acc, err := config.Entities.GetAccountByProviderAndUserId(*config.Context, goauth_entities.GetAccountByProviderAndUserIdParams{Providerid: providerID, Userid: userID})
	if err != nil {
		return nil, err
	}
	if acc.ID == "" {
		return nil, nil
	}
	return &goauth_models.Account{UUID: acc.ID, Provider: acc.Providerid, UserUUID: acc.Userid, Password: acc.Password.String}, nil
}

func GetByUUID(config *goauth_models.Configuration, uuid string) (*goauth_models.Account, error) {
	if uuid == "" {
		return nil, fmt.Errorf("uuid required")
	}
	acc, err := config.Entities.GetAccount(*config.Context, uuid)
	if err != nil {
		return nil, err
	}
	if acc.ID == "" {
		return nil, fmt.Errorf("account not found")
	}
	return &goauth_models.Account{UUID: acc.ID, Provider: acc.Providerid, UserUUID: acc.Userid, Password: acc.Password.String}, nil
}
