package goauth

import (
	"context"
	"fmt"

	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

// SetupConfiguration sets up the configuration for GoAuth
func SetupConfiguration(secret string, entites *goauth_entities.Queries, ctx context.Context, passwordPolicy goauth_models.PasswordPolicy) *goauth_models.Configuration {
	return &goauth_models.Configuration{
		Secret:                   secret,
		SessionDurationInSeconds: 3600,
		Entities:                 entites,
		Context:                  &ctx,
		PasswordPolicy:           &passwordPolicy,
		AllowMultipleAccounts:    true,
	}
}

// SetupProviders sets up the authentication providers
func SetupProviders(providers ...goauth_models.Provider) *goauth_models.ProviderConfig {
	return &goauth_models.ProviderConfig{Providers: providers}
}

// UseProviders returns the provider by name from the configuration
func UseProviders(cfg *goauth_models.ProviderConfig, providerName string) (goauth_models.Provider, error) {
	for _, provider := range cfg.Providers {
		if provider.GetName() == providerName {
			return provider, nil
		}
	}
	return nil, fmt.Errorf("provider %s not found", providerName)
}
