package goauth_models

const (
	EMAIL_AND_PASSWORD = "email_and_password"
	MAGIC_LINK         = "magic_link"
)

// Provider is the interface for an authentication provider
type Provider interface {
	GetName() string
	SignUp(config *Configuration, user *UserUnauthenticated) (*UserAuthenticated, error)
	SignIn(config *Configuration, credentials *Credentials) (*Session, error)
	RecoverPassword(config *Configuration, email string) error
	ResetPassword(config *Configuration, email, token, newPassword string) error
	ValidateEmail(config *Configuration, email, token string) error
	ValidateSession(config *Configuration, token string) (*Session, error)
}

// ProviderConfig is the configuration for an authentication provider
type ProviderConfig struct {
	Providers []Provider
}
