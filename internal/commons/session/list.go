package session

import (
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func ListByUser(config *goauth_models.Configuration, userUUID string) ([]goauth_models.Session, error) {
	ents, err := config.Entities.ListSessionsByUserId(*config.Context, userUUID)
	if err != nil {
		return nil, err
	}
	var out []goauth_models.Session
	for _, s := range ents {
		out = append(out, goauth_models.Session{UUID: s.ID, UserUUID: s.Userid, Token: s.Token})
	}
	return out, nil
}
