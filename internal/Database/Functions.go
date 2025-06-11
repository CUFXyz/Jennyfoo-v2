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
	rows, err := dbi.sqlinstance.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&user.Login, &user.Password, &user.Email); err != nil {
			return &[]models.User{}, err
		}
		users = append(users, user)
	}
	return &users, nil
}

// Getting only one user with specific parameters
func (dbi *DBInstance) GetUser(params models.User) (*models.User, error) {
	var user models.User
	if params.Email == "" || params.Login == "" {
		return nil, fmt.Errorf("empty params not allowed")
	}

	if params.Email != "" {
		query := "SELECT * FROM Users WHERE email = $1"
		row := dbi.sqlinstance.QueryRowx(query, params.Email)
		err := row.Scan(&user.Login, &user.Password, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	if params.Login != "" {
		query := "SELECT * FROM Users WHERE login = $1"
		row := dbi.sqlinstance.QueryRowx(query, params.Login)
		err := row.Scan(&user.Login, &user.Password, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return &user, fmt.Errorf("it's  not to be here")
}

// Sending data about users into db
func (dbi *DBInstance) SendUser(user models.User) error {
	query := "INSERT INTO Users (login, password, email, role) VALUES ($1, $2, $3, $4)"
	_, err := dbi.sqlinstance.Queryx(query, user.Login, user.Password, user.Email, models.USERDEFAULT)
	if err != nil {
		return err
	}

	return nil
}
