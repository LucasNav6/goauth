package goauth_providers

import (
	"github.com/LucasNav6/goauth/internal/providers/magicLink"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

type magicLinkProvider struct{}

func (m *magicLinkProvider) GetName() string {
	return goauth_models.MAGIC_LINK
}

func (m *magicLinkProvider) SignUp(config *goauth_models.Configuration, user *goauth_models.UserUnauthenticated) (*goauth_models.UserAuthenticated, error) {
	// For magic link, signup can simply create the user if not exists
	return nil, nil
}

func (m *magicLinkProvider) SignIn(config *goauth_models.Configuration, credentials *goauth_models.Credentials) (*goauth_models.Session, error) {
	// Not used: magic link sign-in uses email+token API; return error
	return nil, nil
}

func (m *magicLinkProvider) RecoverPassword(config *goauth_models.Configuration, email string) error {
	// Expose RecoverPassword that builds a simple link using a default base URL
	base := "https://example.com/magic"
	return magicLink.RecoverPassword(config, email, base)
}

func (m *magicLinkProvider) ResetPassword(config *goauth_models.Configuration, email string, token string, newPassword string) error {
	// Not applicable for magic link
	return nil
}

func (m *magicLinkProvider) ValidateEmail(config *goauth_models.Configuration, email string, token string) error {
	return magicLink.ValidateEmailWithToken(config, email, token)
}

func (m *magicLinkProvider) ValidateSession(config *goauth_models.Configuration, token string) (*goauth_models.Session, error) {
	return magicLink.ValidateSession(config, token)
}

func MagicLink() goauth_models.Provider {
	return &magicLinkProvider{}
}
