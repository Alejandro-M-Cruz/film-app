package user

import (
	"errors"
	"strconv"
	"time"
)

type UserID int64

func AtoUserID(id string) (UserID, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("invalid user id")
	}

	return UserID(parsedID), nil
}

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
