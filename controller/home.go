package controller

import (
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
)

// Load the routes.
/*func Load() {
	router.Get("/", Index)
}*/

// HomeHandler represents the services required for this controller.
type HomeHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index displays the home page.
func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("home/index")
	if c.Sess.Values["id"] != nil {
		v.Vars["first_name"] = c.Sess.Values["first_name"]
	}

	v.Render(w, r)
}
