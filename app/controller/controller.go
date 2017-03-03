// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/blue-jay/blueshift/app/controller/about"
	"github.com/blue-jay/blueshift/app/controller/debug"
	"github.com/blue-jay/blueshift/app/controller/home"
	"github.com/blue-jay/blueshift/app/controller/login"
	"github.com/blue-jay/blueshift/app/controller/notepad"
	"github.com/blue-jay/blueshift/app/controller/register"
	"github.com/blue-jay/blueshift/app/controller/static"
	"github.com/blue-jay/blueshift/app/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
}
