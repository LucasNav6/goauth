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
func (e *emailAndPasswordProvider) RecoverPassword() (*goauth_models.UserAuthenticated, error) {
	panic("unimplemented")
}

// ResetPassword implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) ResetPassword() (*goauth_models.UserAuthenticated, error) {
	panic("unimplemented")
}

// ValidateEmail implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) ValidateEmail() (*goauth_models.UserAuthenticated, error) {
	panic("unimplemented")
}

// ValidateSession implements [goauth_models.Provider].
func (e *emailAndPasswordProvider) ValidateSession() (*goauth_models.Session, error) {
	panic("unimplemented")
}

func EmailAndPassword() goauth_models.Provider {
	return &emailAndPasswordProvider{}
}
