package permissions

import "film-app/user"

type Policy[T any] interface {
	CanViewAny(usr user.User) bool
	CanView(usr user.User, resource T) bool
	CanCreate(usr user.User) bool
	CanUpdate(usr user.User, resource T) bool
	CanDelete(usr user.User, resource T) bool
}
