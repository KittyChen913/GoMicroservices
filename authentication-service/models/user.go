package models

import (
	"authentication-service/db"
	"context"
	"time"
)

type User struct {
	Id         int
	Email      string
	FirstName  string
	LastName   string
	Password   string
	UserActive int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) QueryUserByEmail(email string) (*User, error) {
	queryStr := `SELECT Id, Email, Password FROM Users WHERE email = $1`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	userRow := db.DbPool.QueryRow(ctx, queryStr, email)
	var user User
	err := userRow.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}
