package main

import (
	"context"
	"fmt"
	"net/http"

	goauth "github.com/LucasNav6/goauth/pkg"
	goauth_entities "github.com/LucasNav6/goauth/pkg/entities"
	goauth_models "github.com/LucasNav6/goauth/pkg/models"
	goauth_providers "github.com/LucasNav6/goauth/pkg/providers"
	"github.com/jackc/pgx/v5"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Connect to the database
		ctx := context.Background()
		conn, err := pgx.Connect(ctx, "postgresql://neondb_owner:npg_qmDEP9Ui8tfZ@ep-steep-paper-ac4c15ul-pooler.sa-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close(ctx)

		queries := goauth_entities.New(conn)

		// Setup GoAuth configuration and providers
		goauthConfig := goauth.SetupConfiguration(queries, ctx, goauth_models.PasswordPolicy{
			MinLength:           8,
			RequireUppercase:    true,
			RequireLowercase:    true,
			RequireNumbers:      true,
			RequireSpecialChars: true,
		})

		// Setup the providers
		goauthProvider := goauth.SetupProviders(
			goauth_providers.EmailAndPassword(),
		)

		// Example of using a provider
		emailAndPasswordProvider, err := goauth.UseProviders(goauthProvider, "email_and_password")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		password := "SecureP@ssw0rd"
		res, err := emailAndPasswordProvider.SignUp(goauthConfig, &goauth_models.UserUnauthenticated{
			Email:    "user@example.com",
			Password: &password,
			Name:     "Nombre",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User created: %+v", res)
	})

	fmt.Println("Server initialize :1111")
	http.ListenAndServe(":1111", nil)
}
