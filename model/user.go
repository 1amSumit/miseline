package model

import (
	"errors"
	"time"

	"sumit.com/mise-link/db"
	"sumit.com/mise-link/utils"
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUser struct {
	Id       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `
    insert into "users" (name, email, password) 
    values ($1, $2, $3) 
    returning id
    `

	hassedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, user.Name, user.Email, hassedPassword).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (loginData *LoginUser) ValidateCredentials() error {
	query := "SELECT id,  password FROM users WHERE email = $1 "
	row := db.DB.QueryRow(query, loginData.Email)

	var reterievedPassword string

	err := row.Scan(&loginData.Id, &reterievedPassword)

	if err != nil {
		return errors.New("credentials Invalid")
	}

	passwordIsvalid := utils.CheckPassword(loginData.Password, reterievedPassword)

	if !passwordIsvalid {
		return errors.New("credentials Invalid")
	}

	return nil
}
