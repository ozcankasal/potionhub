package models

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // Postgres driver
)

type DB struct {
	Conn *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	Conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = Conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: Conn}, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}

// GetUserByID gets user by ID
func (db *DB) GetUserByID(id int64) (*User, error) {
	user := new(User)
	err := db.Conn.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// GetUserByEmail gets user by email
func (db *DB) GetUserByEmail(email string) (*User, error) {
	user := new(User)
	err := db.Conn.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user
func (db *DB) CreateUser(user *User) error {
	err := db.Conn.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Email, user.Password).Scan(&user.ID)
	return err
}

// UpdateUser updates a user
func (db *DB) UpdateUser(user *User) error {
	_, err := db.Conn.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4", user.Username, user.Email, user.Password, user.ID)
	return err
}

// DeleteUser deletes a user by ID
func (db *DB) DeleteUser(id int64) error {
	_, err := db.Conn.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
