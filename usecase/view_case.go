package usecase

import "github.com/blue-jay/blueshift/domain"

// ViewCase represents a service for managing routes.
type ViewCase struct {
	repo domain.ViewRepo
}

// NewViewCase returns the service for managing routes.
func NewViewCase(repo domain.ViewRepo) *ViewCase {
	s := new(ViewCase)
	s.repo = repo
	return s
}

// New returns the service for managing routes.
func (v ViewCase) New(templateList ...string) domain.ViewRepo {
	return v.repo.New(templateList...)
}
