package app

import "github.com/blue-jay/core/view"

// ViewService
type ViewService interface {
	New(name string) *view.Info
}
