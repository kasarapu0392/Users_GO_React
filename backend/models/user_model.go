package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/squirrel"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func CheckUserExists(db *sql.DB, userName string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE user_name = $1`
	var count int
	err := db.QueryRow(query, userName).Scan(&count)
	if err != nil {
		log.Printf("[CheckUserExists] Database error: %v", err)
		return false, err
	}
	return count > 0, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	query := psql.Select("id", "user_name", "email").From("users")
	rows, err := query.RunWith(db).Query()
	if err != nil {
		log.Printf("[GetUsers] Error executing query: %v", err)
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email); err != nil {
			log.Printf("[GetUsers] Error scanning row: %v", err)
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(db *sql.DB, user User) error {
	if user.UserName == "" || user.Email == "" {
		log.Printf("[CreateUser] Missing fields: UserName=%s, Email=%s", user.UserName, user.Email)
		return errors.New("user name and email are required")
	}

	query := psql.Insert("users").
		Columns("user_name", "email").
		Values(user.UserName, user.Email).
		Suffix("RETURNING id")

	err := query.RunWith(db).QueryRow().Scan(&user.ID)
	if err != nil {
		// Handle PostgreSQL-specific "duplicate key" error
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Printf("[CreateUser] Duplicate user error: %v", err)
			return fmt.Errorf("user already exists")
		}

		// Log any other database error
		log.Printf("[CreateUser] Error creating user: %v", err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	log.Printf("[CreateUser] User created with ID: %d", user.ID)
	return nil
}

func UpdateUser(db *sql.DB, user User) error {
	query := psql.Update("users").
		Set("user_name", user.UserName).
		Set("email", user.Email).
		Where(squirrel.Eq{"id": user.ID})

	_, err := query.RunWith(db).Exec()
	if err != nil {
		log.Printf("[UpdateUser] Error updating user: %v", err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := psql.Delete("users").Where(squirrel.Eq{"id": id})

	_, err := query.RunWith(db).Exec()
	if err != nil {
		log.Printf("[DeleteUser] Error deleting user: %v", err)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
