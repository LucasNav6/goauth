package emailAndPassword

import (
	"fmt"

	"github.com/LucasNav6/goauth/internal/commons/account"
	"github.com/LucasNav6/goauth/internal/commons/session"
	"github.com/LucasNav6/goauth/internal/commons/user"
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func SignIn(config *goauth_models.Configuration, credentials *goauth_models.Credentials) (*goauth_models.Session, error) {
	// Check if user exists
	userExist, err := user.GetUserByEmail(config, credentials.Email)
	if err != nil {
		return nil, err
	}
	if userExist.Uuid == "" {
		return nil, fmt.Errorf("user does not exist")
	}

	// Get the account associated with the user
	accountExist, err := account.GetAccountByUserIDAndProvider(config, userExist.Uuid, goauth_models.EMAIL_AND_PASSWORD)
	if err != nil {
		return nil, err
	}
	if accountExist == nil {
		return nil, fmt.Errorf("the user does not have an email and password account")
	}

	// Verify password
	isPasswordValid := utilities.CheckPasswordHash(*credentials.Password, accountExist.Password)
	if err != nil {
		return nil, err
	}
	if !isPasswordValid {
		return nil, fmt.Errorf("invalid password")
	}

	// Create a session for the user
	session, err := session.CreateSession(config, userExist.Uuid, *credentials.UserAgent, *credentials.IP)
	if err != nil {
		return nil, err
	}

	return session, nil
}
