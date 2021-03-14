package model

type Profile struct {
	ID        string `db:"id"`
	Age       uint64 `db:"age"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UserID    string `db:"user_id"`
}
