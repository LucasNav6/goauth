package account

import (
	"fmt"

	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func Delete(config *goauth_models.Configuration, uuid string) error {
	if uuid == "" {
		return fmt.Errorf("uuid required")
	}
	return config.Entities.DeleteAccount(*config.Context, uuid)
}
