package utilities

import (
	"fmt"
	"net/mail"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type PasswordPolicy struct {
	MinLength      int
	RequireUpper   bool
	RequireLower   bool
	RequireDigit   bool
	RequireSpecial bool
}

func IsValidPassword(password string, policy PasswordPolicy) (bool, error) {
	if len(password) < policy.MinLength {
		return false, fmt.Errorf("password must be at least %d characters long", policy.MinLength)
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case (char >= 33 && char <= 47) || (char >= 58 && char <= 64) ||
			(char >= 91 && char <= 96) || (char >= 123 && char <= 126):
			hasSpecial = true
		}
	}

	if policy.RequireUpper && !hasUpper {
		return false, fmt.Errorf("password must contain at least one uppercase letter")
	}
	if policy.RequireLower && !hasLower {
		return false, fmt.Errorf("password must contain at least one lowercase letter")
	}
	if policy.RequireDigit && !hasDigit {
		return false, fmt.Errorf("password must contain at least one digit")
	}
	if policy.RequireSpecial && !hasSpecial {
		return false, fmt.Errorf("password must contain at least one special character")
	}

	return true, nil
}
