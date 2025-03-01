package film

import (
	"film-app/permissions"
	"film-app/user"
)

type Policy struct {
}

func NewPolicy() permissions.Policy[Film] {
	return &Policy{}
}

func (p *Policy) CanViewAny(usr user.User) bool {
	return true
}

func (p *Policy) CanView(usr user.User, film Film) bool {
	return true
}

func (p *Policy) CanCreate(usr user.User) bool {
	return true
}

func (p *Policy) CanUpdate(usr user.User, film Film) bool {
	return usr.ID == film.UserID
}

func (p *Policy) CanDelete(usr user.User, film Film) bool {
	return usr.ID == film.UserID
}
