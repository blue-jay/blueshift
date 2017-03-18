package controller

import (
	"net/http"
	"os"
	"path"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
)

// Load the routes.
/*func Load() {
	// Serve static files
	router.Get("/static/*filepath", Index)
}*/

// StaticHandler represents the services required for this controller.
type StaticHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index maps static files.
func (h *StaticHandler) Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// File path
	path := path.Join(c.Config.Asset.Folder, r.URL.Path[1:])

	// Only serve files
	if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
		http.ServeFile(w, r, path)
		return
	}

	// TODO Figure out how to pass to another handler.
	//status.Error404(w, r)
}
