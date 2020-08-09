package users

import "backend/config"

type User struct {
	ID       int
	Username string
	Avatar   string
	Email    string
	Password string
	FullName string
	Role     string
	Blocked  bool
}

func login(user *User) (User, error) {
	row := config.DB.QueryRow(`SELECT id, username, email, password, avatar, fullname, role
														 FROM users WHERE email = $1`, user.Email)

	var compareUser User
	err := row.Scan(&compareUser.ID, &compareUser.Username, &compareUser.Email,
		&compareUser.Password, &compareUser.Avatar, &compareUser.FullName, &compareUser.Role)
	if err != nil {
		return compareUser, err
	}

	return compareUser, nil
}

func create(email string, password string) error {
	_, err := config.DB.Exec(`INSERT INTO users (email, password)
														VALUES ($1, $2)`, email, password)
	if err != nil {
		return err
	}

	return nil
}
