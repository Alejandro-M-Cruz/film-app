package user

import (
	"strconv"
	"time"
)

type UserID uint

func (id UserID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

type User struct {
	ID        UserID
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
