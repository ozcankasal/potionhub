package models

type User struct {
	ID       int64
	Username string
	Email    string
	Password string // Store the hashed password, not plaintext
}
