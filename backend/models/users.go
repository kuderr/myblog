package models

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	FullName string
	Role     string
	Blocked  bool
}
