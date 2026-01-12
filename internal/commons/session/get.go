package session

import (
	"fmt"

	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func GetByUUID(config *goauth_models.Configuration, uuid string) (*goauth_models.Session, error) {
	if uuid == "" {
		return nil, fmt.Errorf("uuid required")
	}
	s, err := config.Entities.GetSession(*config.Context, uuid)
	if err != nil {
		return nil, err
	}
	if s.ID == "" {
		return nil, fmt.Errorf("session not found")
	}
	return &goauth_models.Session{UUID: s.ID, UserUUID: s.Userid, Token: s.Token}, nil
}

func GetByToken(config *goauth_models.Configuration, token string) (*goauth_models.Session, error) {
	if token == "" {
		return nil, fmt.Errorf("token required")
	}
	s, err := config.Entities.GetSessionByToken(*config.Context, token)
	if err != nil {
		return nil, err
	}
	if s.ID == "" {
		return nil, fmt.Errorf("session not found")
	}
	return &goauth_models.Session{UUID: s.ID, UserUUID: s.Userid, Token: s.Token}, nil
}
