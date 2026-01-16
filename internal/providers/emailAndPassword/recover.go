package emailAndPassword

import (
	"fmt"

	commonsUser "github.com/LucasNav6/goauth/internal/commons/user"
	commonsVerification "github.com/LucasNav6/goauth/internal/commons/verification"
	"github.com/LucasNav6/goauth/internal/utilities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

// RecoverPassword creates a verification token for password recovery. The token
// should be sent to the user's email and is not returned to the caller.
func RecoverPassword(config *goauth_models.Configuration, email string) error {
	if !utilities.IsValidEmail(email) {
		return fmt.Errorf("The email provided is not valid")
	}

	// Ensure user exists; do not reveal existence to caller. If user does not
	// exist, act as if we succeeded to avoid account enumeration.
	u, err := commonsUser.GetByEmail(config, email)
	if err != nil || u == nil || u.Uuid == "" {
		// swallow error and return nil for security: do not indicate existence
		return nil
	}

	// Generate a verification value and persist it; do not expose token to caller.
	token := utilities.GenerateUUID()
	// 15 minutes expiry
	_, err = commonsVerification.Create(config, email, token, int64(15*60))
	if err != nil {
		return fmt.Errorf("failed to create verification: %v", err)
	}

	// Send email if sender is configured
	if config.EmailSender != nil {
		subject := "Password recovery"
		body := "Use the following token to reset your password: " + token
		_ = config.EmailSender(email, subject, body)
	}
	return nil
}
