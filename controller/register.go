package controller

import (
	"errors"
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
	"github.com/blue-jay/blueshift/model/user"

	"github.com/blue-jay/core/passhash"
)

// RegisterHandler represents the services required for this controller.
type RegisterHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index displays the register page.
func (h *RegisterHandler) Index(w http.ResponseWriter, r *http.Request) {
	//c := flight.Context(w, r)
	h.ViewService.SetTemplate("register/index")
	//form.Repopulate(r.Form, v.Vars, "first_name", "last_name", "email")
	h.ViewService.Render(w, r)
}

// Store handles the registration form submission.
func (h *RegisterHandler) Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Validate with required fields
	if !c.FormValid("first_name", "last_name", "email", "password", "password_verify") {
		h.Index(w, r)
		return
	}

	// Get form values
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")

	// Validate passwords
	if r.FormValue("password") != r.FormValue("password_verify") {
		c.FlashError(errors.New("Passwords do not match."))
		h.Index(w, r)
		return
	}

	// Hash password
	password, errp := passhash.HashString(r.FormValue("password"))

	// If password hashing failed
	if errp != nil {
		c.FlashErrorGeneric(errp)
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	// Get database result
	_, noRows, err := user.ByEmail(c.DB, email)

	if noRows { // If success (no user exists with that email)
		_, err = user.Create(c.DB, firstName, lastName, email, password)
		// Will only error if there is a problem with the query
		if err != nil {
			c.FlashErrorGeneric(err)
		} else {
			c.FlashSuccess("Account created successfully for: " + email)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	} else if err != nil { // Catch all other errors
		c.FlashErrorGeneric(err)
	} else { // Else the user already exists
		c.FlashError(errors.New("Account already exists for: " + email))
	}

	// Display the page
	h.Index(w, r)
}
