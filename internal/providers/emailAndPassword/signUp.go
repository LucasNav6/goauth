package emailAndPassword

import (
	"fmt"

	"github.com/LucasNav6/goauth/internal/commons/account"
	commons "github.com/LucasNav6/goauth/internal/commons/user"
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func SignUp(config *goauth_models.Configuration, user *goauth_models.UserUnauthenticated) (*goauth_models.UserAuthenticated, error) {
	// Validate user input (email, password)
	if !utilities.IsValidEmail(user.Email) {
		err := fmt.Errorf("invalid email format")
		return nil, err
	}

	// Validate password policy
	if user.Password != nil {
		isValid, err := utilities.IsValidPassword(*user.Password, utilities.PasswordPolicy{
			MinLength:      config.PasswordPolicy.MinLength,
			RequireUpper:   config.PasswordPolicy.RequireUppercase,
			RequireLower:   config.PasswordPolicy.RequireLowercase,
			RequireDigit:   config.PasswordPolicy.RequireNumbers,
			RequireSpecial: config.PasswordPolicy.RequireSpecialChars,
		})
		if err != nil {
			return nil, fmt.Errorf("error validating password: %v", err)
		}
		if !isValid {
			return nil, fmt.Errorf("password does not meet the policy requirements")
		}
	} else {
		return nil, fmt.Errorf("password is required")
	}

	// Create a new user
	newUser, err := commons.Create(config, *user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	// Create an account linked to the new user
	_, err = account.CreateWithPassword(config, newUser.Uuid, *user.Password)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %v", err)
	}

	return newUser, nil
}
