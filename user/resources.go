package user

type Resource struct {
    ID       UserID `json:"id"`
    Username string `json:"username"`
}

func NewResource(user User) Resource {
    return Resource{
        ID:       user.ID,
        Username: user.Username,
    }
}
