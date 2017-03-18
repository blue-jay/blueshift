package controller

import (
	"net/http"
	"net/http/pprof"

	"github.com/blue-jay/blueshift/domain"
	"github.com/husobee/vestigo"
)

// Load the routes.
/*func Load() {
	// Enable Pprof
	router.Get("/debug/pprof/", Index, acl.DisallowAnon)
	router.Get("/debug/pprof/:pprof", Profile, acl.DisallowAnon)
}*/

// DebugHandler represents the services required for this controller.
type DebugHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index shows the profile index.
func (h *DebugHandler) Index(w http.ResponseWriter, r *http.Request) {
	pprof.Index(w, r)
}

// Profile shows the individual profiles.
func (h *DebugHandler) Profile(w http.ResponseWriter, r *http.Request) {
	switch vestigo.Param(r, "pprof") {
	case "cmdline":
		pprof.Cmdline(w, r)
	case "profile":
		pprof.Profile(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	case "trace":
		pprof.Trace(w, r)
	default:
		h.Index(w, r)
	}
}
