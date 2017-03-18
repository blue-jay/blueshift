package usecase

import (
	"net/http"

	"github.com/blue-jay/blueshift/domain"
)

// RouterCase represents a service for managing routes.
type RouterCase struct {
	routerRepo domain.RouterRepo
}

// NewRouterCase returns the service for managing routes.
func NewRouterCase(repo domain.RouterRepo) *RouterCase {
	s := new(RouterCase)
	s.routerRepo = repo
	return s
}

func (s *RouterCase) Get(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	s.routerRepo.Get(path, handler, acl...)
}

func (s *RouterCase) Delete(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	s.routerRepo.Delete(path, handler, acl...)
}

func (s *RouterCase) Post(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	s.routerRepo.Post(path, handler, acl...)
}

func (s *RouterCase) Patch(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	s.routerRepo.Patch(path, handler, acl...)
}

func (s *RouterCase) Put(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	s.routerRepo.Put(path, handler, acl...)
}
