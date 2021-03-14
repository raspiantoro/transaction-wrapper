package model

type User struct {
	ID       string `db:"id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}
