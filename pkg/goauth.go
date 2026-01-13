package goauth

import (
	"context"
	"fmt"

	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

func SetupConfiguration(options ...func(*goauth_models.Configuration)) *goauth_models.Configuration {
	cfg := &goauth_models.Configuration{}

	for _, option := range options {
		option(cfg)
	}

	return cfg
}

func SetupSecret(secret string) func(*goauth_models.Configuration) {
	return func(cfg *goauth_models.Configuration) {
		cfg.Secret = secret
	}
}

func SetupDatabase(ctx context.Context, entities *goauth_entities.Queries) func(*goauth_models.Configuration) {
	return func(cfg *goauth_models.Configuration) {
		cfg.Context = &ctx
		cfg.Entities = entities
	}
}

func SetupSession(durationInSeconds int) func(*goauth_models.Configuration) {
	return func(cfg *goauth_models.Configuration) {
		cfg.SessionDurationInSeconds = int64(durationInSeconds)
	}
}

func PasswordPolicy(policy *goauth_models.PasswordPolicy) func(*goauth_models.Configuration) {
	return func(cfg *goauth_models.Configuration) {
		cfg.PasswordPolicy = policy
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
