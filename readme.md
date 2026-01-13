# GoAuth
GoAuth is a minimalistic and highly customizable Go authentication library that simplifies authentication for your applications. It's highly scalable, enabling developers to integrate only the providers they need. Designed to ease development at the code level, database integration, and security.

## Features
- Highly configurable setup. Enable only the providers you'll use in your project for optimal performance and simplicity.

## Instalation
```bash
go get -u github.com/LucasNav6/goauth@vX.X.X
```
## Implement GoAuth in Your Project

### Goauth configuration
First, you need to configure the GoAuth project. In this configuration, you define how the application will handle user management with different providers, including settings for cookies, session duration, validation requirements, and more.

```go
import (
    goauth "github.com/LucasNav6/goauth/pkg"
    goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

goauthConfig := goauth.SetupConfiguration(
    // Always required
    goauth.SetupSecret("YOUR-SECURITY-SECRET")
    goauth.SetupDatabase(context, queries)
        
    // Optional
    goauth.SetupSession(
        session_duration_in_second // Integer
    )

    // Depending on the provider
    // Email and password
    goauth.PasswordPolicy(&goauth_models.PasswordPolicy{
        MinLength           int
        RequireUppercase    bool
        RequireLowercase    bool
        RequireNumbers      bool
        RequireSpecialChars bool
    })

    // Other functions...
)
```