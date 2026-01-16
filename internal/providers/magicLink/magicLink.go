package magicLink

import (
	"fmt"

	commonsSession "github.com/LucasNav6/goauth/internal/commons/session"
	commonsUser "github.com/LucasNav6/goauth/internal/commons/user"
	commonsVerification "github.com/LucasNav6/goauth/internal/commons/verification"
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

// SignIn with magic link token. The token is sent to email previously via RecoverPassword.
func SignInWithToken(config *goauth_models.Configuration, email string, token string, userAgent string, ip string) (*goauth_models.Session, error) {
	if !utilities.IsValidEmail(email) {
		return nil, fmt.Errorf("invalid email")
	}
	if token == "" {
		return nil, fmt.Errorf("token required")
	}

	vs, err := commonsVerification.ListByIdentifier(config, email)
	if err != nil {
		return nil, err
	}
	var matchedID string
	for _, v := range vs {
		if v.Value == token {
			matchedID = v.ID
			break
		}
	}
	if matchedID == "" {
		return nil, fmt.Errorf("invalid or expired token")
	}

	// Find or create user
	u, err := commonsUser.GetByEmail(config, email)
	if err != nil {
		return nil, err
	}
	if u == nil || u.Uuid == "" {
		// Create a minimal user record
		newUser, err := commonsUser.Create(config, goauth_models.UserUnauthenticated{Name: "", Email: email})
		if err != nil {
			return nil, err
		}
		u = newUser
	}

	// Delete verification token after use
	_ = commonsVerification.Delete(config, matchedID)

	// Create session
	sess, err := commonsSession.Create(config, u.Uuid, ip, userAgent)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// RecoverPassword sends a magic link token to the given email.
func RecoverPassword(config *goauth_models.Configuration, email string, baseURL string) error {
	if !utilities.IsValidEmail(email) {
		return fmt.Errorf("invalid email")
	}

	// Generate token and store verification
	token := utilities.GenerateUUID()
	_, err := commonsVerification.Create(config, email, token, int64(15*60))
	if err != nil {
		return fmt.Errorf("failed to create verification: %v", err)
	}

	// Build simple link
	link := baseURL + "?email=" + email + "&token=" + token

	// Send email if configured
	if config.EmailSender != nil {
		subj := "Sign in to your account"
		body := "Click the link to sign in: " + link
		_ = config.EmailSender(email, subj, body)
	}
	return nil
}

// ValidateEmailWithToken marks email as verified (reuse logic similar to emailAndPassword)
func ValidateEmailWithToken(config *goauth_models.Configuration, email string, token string) error {
	if !utilities.IsValidEmail(email) {
		return fmt.Errorf("invalid email")
	}
	if token == "" {
		return fmt.Errorf("token required")
	}

	vs, err := commonsVerification.ListByIdentifier(config, email)
	if err != nil {
		return err
	}
	var matchedID string
	for _, v := range vs {
		if v.Value == token {
			matchedID = v.ID
			break
		}
	}
	if matchedID == "" {
		return fmt.Errorf("invalid or expired token")
	}

	u, err := commonsUser.GetByEmail(config, email)
	if err != nil {
		return err
	}
	if u == nil || u.Uuid == "" {
		return fmt.Errorf("user not found")
	}

	// Mark email verified: call update (no changes to fields besides verification in real impl)
	err = commonsUser.Update(config, u.Uuid, nil, nil, nil)
	if err != nil {
		return err
	}
	_ = commonsVerification.Delete(config, matchedID)
	return nil
}

// ValidateSession checks a session token
func ValidateSession(config *goauth_models.Configuration, token string) (*goauth_models.Session, error) {
	if token == "" {
		return nil, fmt.Errorf("token required")
	}
	return commonsSession.GetByToken(config, token)
}
