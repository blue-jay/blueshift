package controller

import (
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
)

// Load the routes.
/*func Load() {
	router.Get("/about", Index)
}*/

// AboutHandler represents the services required for this controller.
type AboutHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index displays the About page.
func (h *AboutHandler) Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	c.View.New("about/index").Render(w, r)
}
