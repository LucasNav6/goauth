package goauth_models

const (
	EMAIL_AND_PASSWORD = "email_and_password"
	MAGIC_LINK         = "magic_link"
)

// Provider is the interface for an authentication provider
type Provider interface {
	GetName() string                                                         // Returns the name of the provider
	SignUp(*Configuration, *UserUnauthenticated) (*UserAuthenticated, error) // Creates a new user
	SignIn(*Configuration, *Credentials) (*Session, error)                   // Signs in an existing user
	ResetPassword() (*UserAuthenticated, error)                              // Resets the user's password
	RecoverPassword() (*UserAuthenticated, error)                            // Sends a password recovery email
	ValidateEmail() (*UserAuthenticated, error)                              // Sends an email validation link
	ValidateSession() (*Session, error)                                      // Validates an existing session
}

// ProviderConfig is the configuration for an authentication provider
type ProviderConfig struct {
	Providers []Provider
}
