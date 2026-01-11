package goauth_models

type UserUnauthenticated struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password,omitempty"` // Not all providers require a password
	Image    *string `json:"image,omitempty"`    // Optional profile image URL
}

type Credentials struct {
	Email     string
	Password  *string
	UserAgent *string
	IP        *string
}

type UserAuthenticated struct {
	Uuid  string
	Name  string
	Email string
}
