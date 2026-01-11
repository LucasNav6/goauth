package goauth_models

import (
	"context"

	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
)

type PasswordPolicy struct {
	MinLength           int
	RequireUppercase    bool
	RequireLowercase    bool
	RequireNumbers      bool
	RequireSpecialChars bool
}

type Configuration struct {
	// JWT
	Secret string

	// Session
	SessionDurationInSeconds int64

	// Database Entities
	Entities *goauth_entities.Queries
	Context  *context.Context

	// Password Policy
	PasswordPolicy *PasswordPolicy

	// Account
	AllowMultipleAccounts bool
}
