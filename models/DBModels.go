package models

const (
	USERDEFAULT = "user"
)

type User struct {
	Login    string `db:"login"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Role     string `db:"role"`
}
