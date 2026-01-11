# GoAuth
GoAuth is a minimalistic and highly customizable Go authentication library that simplifies authentication for your applications. It's highly scalable, enabling developers to integrate only the providers they need. Designed to ease development at the code level, database integration, and security.

## Features
- Highly configurable setup. Enable only the providers you'll use in your project for optimal performance and simplicity.

## Instalation
```bash
go get -u github.com/LucasNav6/goauth
```
## Implement GoAuth in Your Project

### Configure the GoAuth Providers
First, you need to configure which providers are available in your project. This returns a standard `goauthCfg` object. It is not recommended to change this before setting up providers.

```go
goauthCfg := goauth.SetupProviders(
    goauth.EmailAndPassword(),
    goauth.MagicLink(),
    // Other providers...
)
```

### Use a GoAuth Provider
Next, you can use only the providers that you configured in the previous step. If the specified provider does not exist, an error will be returned.

```go
goauthProvider, err := goauth.UseProviders(goauthCfg, "email_and_password")
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```
