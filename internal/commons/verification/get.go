package verification

import (
	"fmt"

	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func GetByID(config *goauth_models.Configuration, id string) (*goauth_entities.Verification, error) {
	if id == "" {
		return nil, fmt.Errorf("id required")
	}
	v, err := config.Entities.GetVerification(*config.Context, id)
	if err != nil {
		return nil, err
	}
	if v.ID == "" {
		return nil, fmt.Errorf("verification not found")
	}
	return &v, nil
}

func ListByIdentifier(config *goauth_models.Configuration, identifier string) ([]goauth_entities.Verification, error) {
	if identifier == "" {
		return nil, fmt.Errorf("identifier required")
	}
	vs, err := config.Entities.GetVerificationByIdentifier(*config.Context, identifier)
	if err != nil {
		return nil, err
	}
	return vs, nil
}
