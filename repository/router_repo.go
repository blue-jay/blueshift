package repository

import (
	"net/http"

	"github.com/blue-jay/core/router"
	"github.com/justinas/alice"
)

// RouterRepo represents a service for managing routes.
type RouterRepo struct {
}

// NewRouterRepo returns the service for managing routes.
/*func NewRouterRepo(repo domain.RouterRepo) *RouterRepo {
	s := new(RouterRepo)
	s.routerRepo = repo
	return s
}*/

func (s *RouterRepo) Get(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	c := []alice.Constructor{}
	for _, v := range acl {
		c = append(c, v)
	}

	router.Get(path, handler, c...)
}

func (s *RouterRepo) Delete(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	//router.Delete(path, handler, acl...)
}

func (s *RouterRepo) Post(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	//router.Post(path, handler, acl...)
}

func (s *RouterRepo) Patch(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	//router.Patch(path, handler, acl...)
}

func (s *RouterRepo) Put(path string, handler http.HandlerFunc, acl ...func(http.Handler) http.Handler) {
	//router.Put(path, handler, acl...)
}
