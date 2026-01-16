# GoAuth

GoAuth is a minimalistic and highly customizable Go authentication library that simplifies authentication for your applications. It allows you to integrate only the providers you need, making it scalable for both small and large projects. Designed to streamline development at the code level, database integration, and security.

## Features
- Highly configurable setup: enable only the providers you need for optimal performance and simplicity.
- Scalability: suitable for projects of any size.
- Robust security: includes password policies and session management.

## Installation
```bash
go get -u github.com/LucasNav6/goauth@vX.X.X
```

## Implementing GoAuth in Your Project

### GoAuth Configuration
First, configure GoAuth by defining how the application will handle user management with different providers. This includes settings for cookies, session duration, validation requirements, and more.

```go
import (
    "context"
    "os"
    "time"

    goauth "github.com/LucasNav6/goauth/pkg"
    goauth_models "github.com/LucasNav6/goauth/pkg/models"
)

// Basic GoAuth configuration
goauthConfig := goauth.SetupConfiguration(
    // Always required
    goauth.SetupSecret("YOUR-SECURITY-SECRET"),
    goauth.SetupDatabase(context, queries),

    // Optional
    goauth.SetupSession(
        session_duration_in_seconds, // Session duration in seconds
    ),
)
```

#### Complete Example:
```go
// Example configuration
ctx := context.Background()
queries := &YourDatabaseQueries{}

goauthConfig := goauth.SetupConfiguration(
    goauth.SetupSecret(os.Getenv("GOAUTH_SECRET")),
    goauth.SetupDatabase(ctx, queries),
    goauth.SetupSession(int(time.Hour.Seconds())),
)
```

With this configuration, you can customize GoAuth to meet the specific needs of your project.