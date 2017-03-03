package mock

import (
	"github.com/blue-jay/blueshift/app"
)

// UserService represents a mock implementation of app.UserService.
type UserService struct {
	UserByEmailFn      func(email string) (*app.User, error)
	UserByEmailInvoked bool

	UserCreateFn      func() error
	UserCreateInvoked bool

	// additional function implementations...
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) UserByEmail(email string) (*app.User, error) {
	s.UserByEmailInvoked = true
	return s.UserByEmailFn(email)
}

// CreateUser invokes the mock implementation and marks the function as invoked.
func (s *UserService) CreateUser(u *app.User) error {
	s.UserCreateInvoked = true
	return s.UserCreateFn()
}
