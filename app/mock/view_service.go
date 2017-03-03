package mock

import "github.com/blue-jay/core/view"

// ViewService represents a mock implementation of app.ViewService.
type ViewService struct {
	NewFn      func(name string) *view.Info
	NewInvoked bool
}

// User invokes the mock implementation and marks the function as invoked.
func (s *ViewService) New(name string) *view.Info {
	s.NewInvoked = true
	return s.NewFn(name)
}
