package controller

import (
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
)

// Load the routes.
/*func Load() {
	router.MethodNotAllowed(Error405)
	router.NotFound(Error404)
}*/

// StatusHandler represents the services required for this controller.
type StatusHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Error404 - Page Not Found.
func (h *StatusHandler) Error404(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	w.WriteHeader(http.StatusNotFound)
	v := c.View.New("status/index")
	v.Vars["title"] = "404 Not Found"
	v.Vars["message"] = "Page could not be found."
	v.Render(w, r)
}

// Error405 - Method Not Allowed.
func Error405(allowedMethods string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		v := c.View.New("status/index")
		v.Vars["title"] = "405 Method Not Allowed"
		v.Vars["message"] = "Method is not allowed."
		v.Render(w, r)
	}
}

// Error500 - Internal Server Error.
func (h *StatusHandler) Error500(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	w.WriteHeader(http.StatusInternalServerError)
	v := c.View.New("status/index")
	v.Vars["title"] = "500 Internal Server Error"
	v.Vars["message"] = "An internal server error occurred."
	v.Render(w, r)
}

// Error501 - Not Implemented.
func (h *StatusHandler) Error501(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	w.WriteHeader(http.StatusNotImplemented)
	v := c.View.New("status/index")
	v.Vars["title"] = "501 Not Implemented"
	v.Vars["message"] = "Page is not yet implemented."
	v.Render(w, r)
}

// StatusInvalidToken shows a page in response to CSRF attacks.
func StatusInvalidToken(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	w.WriteHeader(http.StatusForbidden)
	v := c.View.New("status/index")
	v.Vars["title"] = "Invalid Token"
	v.Vars["message"] = `Your token <strong>expired</strong>,
		click <a href="javascript:void(0)" onclick="location.replace(document.referrer)">here</a>
		to try again.`
	v.Render(w, r)
}
