package services

import (
	"database/sql"
	"errors"

	"github.com/ozcankasal/potionhub/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *models.DB
}

func NewUserService(db *models.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) GetUserByID(id int64) (*models.User, error) {
	user := new(models.User)
	err := us.db.Conn.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := us.db.Conn.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (us *UserService) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return us.db.CreateUser(user)
}

func (us *UserService) UpdateUser(user *models.User) error {
	_, err := us.db.Conn.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4", user.Username, user.Email, user.Password, user.ID)
	return err
}

func (us *UserService) DeleteUser(id int64) error {
	_, err := us.db.Conn.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (us *UserService) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := us.db.Conn.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
