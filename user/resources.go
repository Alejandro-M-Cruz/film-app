package user

import "film-app/models"

type Resource struct {
    ID       models.UserID `json:"id"`
    Username string        `json:"username"`
}

func NewResource(user models.User) Resource {
    return Resource{
        ID:       user.ID,
        Username: user.Username,
    }
}
