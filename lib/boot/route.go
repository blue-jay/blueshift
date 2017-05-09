package boot

import (
	"net/http"

	"github.com/blue-jay/blueshift/controller"
	"github.com/blue-jay/blueshift/middleware/acl"

	"github.com/blue-jay/core/router"
)

// LoadRoutes returns a handler with all the routes.
func (s *Service) LoadRoutes() http.Handler {
	// Create the mux.
	//h := http.NewServeMux()
	//h := vestigo.NewRouter()
	//h := router.

	// Register the pages.
	s.AddStatic()
	//s.AddLogin(h)
	s.AddRegister()

	// Return the handler.
	return router.Instance()
}

// AddStatic registers the static handlers.
func (s *Service) AddStatic() {
	// Create handler.
	h := new(controller.StaticHandler)

	// Load routes.
	s.RouterService.Get("/static/*filepath", h.Index)

}

// AddLogin registers the login handlers.
func (s *Service) AddLogin() {
	// Create handler.
	h := new(controller.LoginHandler)

	// Assign services.
	h.UserService = s.UserService
	h.ViewService = s.ViewService

	// Load routes.
	//mux.HandleFunc("/", h.Index)
}

// AddRegister registers the register handlers.
func (s *Service) AddRegister() {
	// Create handler.
	h := new(controller.RegisterHandler)

	// Assign services.
	h.UserService = s.UserService
	h.ViewService = s.ViewService

	// Load routes.
	s.RouterService.Get("/register", h.Index, acl.DisallowAuth)
	s.RouterService.Post("/register", h.Store, acl.DisallowAuth)

}
