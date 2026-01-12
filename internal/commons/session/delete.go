package session

import (
	"fmt"

	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func Delete(config *goauth_models.Configuration, uuid string) error {
	if uuid == "" {
		return fmt.Errorf("uuid required")
	}
	return config.Entities.DeleteSession(*config.Context, uuid)
}
