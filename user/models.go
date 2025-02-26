package user

import "time"

type UserID uint

type User struct {
	ID        UserID
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
