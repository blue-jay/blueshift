package boot

import (
	"net/http"

	"github.com/blue-jay/blueshift/lib/env"
)

// ServicesAndRoutes returns an HTTP handler after registering the services
// and loading the routes.
func ServicesAndRoutes(config *env.Info) http.Handler {
	return RegisterServices(config).LoadRoutes()
}
