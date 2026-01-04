package models

import "github.com/jackc/pgx/v5/pgtype"

type ICreateUser struct {
	Email    pgtype.Text
	Name     pgtype.Text
	Image    pgtype.Text
	Password string
}
