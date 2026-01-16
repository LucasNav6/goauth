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
		return nil, fmt.Errorf("The email provided is not valid")
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
			return nil, fmt.Errorf("Failed to validate password: %v", err)
		}
		if !isValid {
			return nil, fmt.Errorf("The password does not meet the policy requirements")
		}
	} else {
		return nil, fmt.Errorf("The password is required")
	}

	// Create a new user
	newUser, err := commons.Create(config, *user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user")
	}

	// Create an account linked to the new user
	_, err = account.CreateWithPassword(config, newUser.Uuid, *user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create account")
	}

	return newUser, nil
}
