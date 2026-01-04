package models

type UserUnauthenticated struct {
	Name     string
	Email    string
	Password string
	Image    string `json:",omitempty"`
}

type UserAuthenticated struct {
	ID    string
	Name  string
	Email string
	Image string `json:",omitempty"`
}
