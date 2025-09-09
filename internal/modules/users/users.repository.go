// internal/modules/users/repository.go
package users

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}

func GetAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func CreateNewUser(req CreateUserRequest) (*User, error) {
	var user User
	err := db.QueryRow(
		"INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, name, email, created_at, updated_at",
		req.Name, req.Email, time.Now(), time.Now(),
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateExistingUser(id int, req UpdateUserRequest) (*User, error) {
	var user User
	err := db.QueryRow(
		"UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4 RETURNING id, name, email, created_at, updated_at",
		req.Name, req.Email, time.Now(), id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
