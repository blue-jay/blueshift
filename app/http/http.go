package http

import (
	"net/http"

	"github.com/blue-jay/blueshift/app"
)

type Handler struct {
	UserService app.UserService
	ViewService app.ViewService
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle request
	h.UserService.UserByEmail("example@example.com")
	h.ViewService.New("home/index")
	//v.Render(w, r)
}
