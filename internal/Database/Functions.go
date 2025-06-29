package database

import (
	"fmt"
	"v2/models"

	_ "github.com/lib/pq"
)

// Getting all users from database
func (dbi *DBInstance) GetUsers() (*[]models.User, error) {
	var users []models.User
	var user models.User
	query := "SELECT * FROM Users"
	rows, err := dbi.Sqlinstance.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.Login, &user.Password, &user.Email, &user.Role); err != nil {
			return &[]models.User{}, err
		}
		users = append(users, user)
	}
	return &users, nil
}

// Getting only one user with specific parameters
func (dbi *DBInstance) GetUser(params *models.User) (*models.User, error) {
	var user models.User

	if params.Email != "" {
		query := "SELECT * FROM Users WHERE email = $1"
		row := dbi.Sqlinstance.QueryRowx(query, params.Email)
		err := row.Scan(&user.Login, &user.Password, &user.Email, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("email: %v", err)
		}
		return &user, nil
	}

	if params.Login != "" {
		query := "SELECT * FROM Users WHERE login = $1"
		row := dbi.Sqlinstance.QueryRowx(query, params.Login)
		err := row.Scan(&user.Login, &user.Password, &user.Email, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("login: %v", err)
		}
		return &user, nil
	}

	if params.Password != "" {
		query := "SELECT * FROM Users WHERE password = $1"
		row := dbi.Sqlinstance.QueryRowx(query, params.Password)
		err := row.Scan(&user.Login, &user.Password, &user.Email, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("password: %v", err)
		}
		return &user, nil
	}

	return &user, fmt.Errorf("it's not to be here")
}

// Sending data about users into db
func (dbi *DBInstance) SendUser(user models.User) error {
	query := "INSERT INTO Users (login, password, email, role) VALUES ($1, $2, $3, $4)"
	_, err := dbi.Sqlinstance.Queryx(query, user.Login, user.Password, user.Email, models.USERDEFAULT)
	if err != nil {
		return err
	}

	return nil
}

func (dbi *DBInstance) PromoteUser(user models.User, role string) error {

	done := false

	if user.Login != "" {
		query := "UPDATE Users SET role = $1 WHERE login = $2"
		_, err := dbi.Sqlinstance.Queryx(query, role, user.Login)
		if err != nil {
			return err
		}
		done = true
	}

	if user.Email != "" && !done {
		query := "UPDATE Users SET role = $1 WHERE email = $2"
		_, err := dbi.Sqlinstance.Queryx(query, role, user.Email)
		if err != nil {
			return err
		}
		done = true
	}

	return nil
}
