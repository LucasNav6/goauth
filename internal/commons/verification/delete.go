package verification

import (
	"fmt"

	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func Delete(config *goauth_models.Configuration, id string) error {
	if id == "" {
		return fmt.Errorf("id required")
	}
	return config.Entities.DeleteVerification(*config.Context, id)
}
