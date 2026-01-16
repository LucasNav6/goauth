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
	userExist, err := user.GetByEmail(config, credentials.Email)
	if err != nil {
		// Do not reveal whether the email exists
		return nil, fmt.Errorf("The credentials provided are invalid")
	}
	if userExist.Uuid == "" {
		return nil, fmt.Errorf("The credentials provided are invalid")
	}

	// Get the account associated with the user
	accountExist, err := account.GetByUserAndProvider(config, userExist.Uuid, goauth_models.EMAIL_AND_PASSWORD)
	if err != nil {
		return nil, err
	}
	if accountExist == nil {
		return nil, fmt.Errorf("The credentials provided are invalid")
	}

	// Verify password
	isPasswordValid := utilities.CheckPasswordHash(*credentials.Password, accountExist.Password)
	if !isPasswordValid {
		// Keep message generic to avoid confirming account existence
		return nil, fmt.Errorf("The credentials provided are invalid")
	}

	// Create a session for the user. The returned session contains a token which
	// server handlers should set as an HTTP-only cookie instead of returning it
	// in a JSON body.
	session, err := session.Create(config, userExist.Uuid, *credentials.UserAgent, *credentials.IP)
	if err != nil {
		return nil, err
	}

	return session, nil
}
