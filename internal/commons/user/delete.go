package user

import (
    "fmt"

    goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

// Delete removes a user by uuid.
func Delete(config *goauth_models.Configuration, uuid string) error {
    if uuid == "" {
        return fmt.Errorf("uuid required")
    }
    return config.Entities.DeleteUser(*config.Context, uuid)
}
