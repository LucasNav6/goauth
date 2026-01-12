package goauth_providers

import (
	"github.com/LucasNav6/goauth/internal/providers/emailAndPassword"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

type emailAndPasswordProvider struct{}

// GetName implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) GetName() string {
	return goauth_models.EMAIL_AND_PASSWORD
}

// SignUp implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) SignUp(config *goauth_models.Configuration, user *goauth_models.UserUnauthenticated) (*goauth_models.UserAuthenticated, error) {
	return emailAndPassword.SignUp(config, user)
}

// SignIn implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) SignIn(config *goauth_models.Configuration, credentials *goauth_models.Credentials) (*goauth_models.Session, error) {
	return emailAndPassword.SignIn(config, credentials)
}

// RecoverPassword implements [goauth_models.Provider].
// Note: provider-level methods are thin wrappers; application code should pass required params.
func (e *emailAndPasswordProvider) RecoverPassword(config *goauth_models.Configuration, email string) error {
	return emailAndPassword.RecoverPassword(config, email)
}

// ResetPassword implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) ResetPassword(config *goauth_models.Configuration, email string, token string, newPassword string) error {
	return emailAndPassword.ResetPasswordWithToken(config, email, token, newPassword)
}

// ValidateEmail implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) ValidateEmail(config *goauth_models.Configuration, email string, token string) error {
	return emailAndPassword.ValidateEmailWithToken(config, email, token)
}

// ValidateSession implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) ValidateSession(config *goauth_models.Configuration, token string) (*goauth_models.Session, error) {
	return emailAndPassword.ValidateSession(config, token)
}

func EmailAndPassword() goauth_models.Provider {
	return &emailAndPasswordProvider{}
}
