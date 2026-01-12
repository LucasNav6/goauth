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
		return fmt.Errorf("invalid email")
	}

	// Ensure user exists
	u, err := commonsUser.GetByEmail(config, email)
	if err != nil {
		return err
	}
	if u == nil || u.Uuid == "" {
		return fmt.Errorf("user not found")
	}

	// Generate a verification value and persist it; do not expose token to caller.
	token := utilities.GenerateUUID()
	// 15 minutes expiry
	_, err = commonsVerification.Create(config, email, token, int64(15*60))
	if err != nil {
		return fmt.Errorf("failed to create verification: %v", err)
	}

	// In a real system we'd send the token by email. Here we do not return it.
	return nil
}
