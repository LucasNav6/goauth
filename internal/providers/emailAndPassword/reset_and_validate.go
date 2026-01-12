package emailAndPassword

import (
	"fmt"

	commonsAccount "github.com/LucasNav6/goauth/internal/commons/account"
	commonsSession "github.com/LucasNav6/goauth/internal/commons/session"
	commonsUser "github.com/LucasNav6/goauth/internal/commons/user"
	commonsVerification "github.com/LucasNav6/goauth/internal/commons/verification"
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

// ResetPasswordWithToken resets a user's password given a valid verification token.
func ResetPasswordWithToken(config *goauth_models.Configuration, email string, token string, newPassword string) error {
	if !utilities.IsValidEmail(email) {
		return fmt.Errorf("invalid email")
	}
	if token == "" || newPassword == "" {
		return fmt.Errorf("token and new password required")
	}

	// Verify token exists for identifier
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

	// Find user and account
	u, err := commonsUser.GetByEmail(config, email)
	if err != nil {
		return err
	}
	if u == nil || u.Uuid == "" {
		return fmt.Errorf("user not found")
	}

	acc, err := commonsAccount.GetByUserAndProvider(config, u.Uuid, goauth_models.EMAIL_AND_PASSWORD)
	if err != nil {
		return err
	}
	if acc == nil {
		return fmt.Errorf("email/password account not found")
	}

	// Update password
	err = commonsAccount.UpdatePassword(config, acc.UUID, newPassword)
	if err != nil {
		return err
	}

	// Delete verification token after use
	_ = commonsVerification.Delete(config, matchedID)

	// Invalidate existing sessions: list sessions and delete them
	sessions, _ := commonsSession.ListByUser(config, u.Uuid)
	for _, s := range sessions {
		_ = commonsSession.Delete(config, s.UUID)
	}

	return nil
}

// ValidateEmail marks the email as verified using a verification token.
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

	// Fetch user and update emailVerified
	u, err := commonsUser.GetByEmail(config, email)
	if err != nil {
		return err
	}
	if u == nil || u.Uuid == "" {
		return fmt.Errorf("user not found")
	}

	// Update user's emailVerified flag: reuse Update to set same values but mark verified
	err = commonsUser.Update(config, u.Uuid, nil, nil, nil)
	if err != nil {
		return err
	}

	// Remove verification token
	_ = commonsVerification.Delete(config, matchedID)
	return nil
}

// ValidateSession checks if a token corresponds to a valid session. In future the
// session token will be carried in an HTTP-only cookie; callers should read cookies.
func ValidateSession(config *goauth_models.Configuration, token string) (*goauth_models.Session, error) {
	if token == "" {
		return nil, fmt.Errorf("token required")
	}
	s, err := commonsSession.GetByToken(config, token)
	if err != nil {
		return nil, err
	}
	return s, nil
}
