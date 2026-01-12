package account

import (
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func ListByUser(config *goauth_models.Configuration, userUUID string) ([]goauth_models.Account, error) {
	ents, err := config.Entities.ListAccountsByUserId(*config.Context, userUUID)
	if err != nil {
		return nil, err
	}
	var out []goauth_models.Account
	for _, a := range ents {
		out = append(out, goauth_models.Account{UUID: a.ID, Provider: a.Providerid, UserUUID: a.Userid, Password: a.Password.String})
	}
	return out, nil
}
