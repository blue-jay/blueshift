package controller

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"

	"github.com/blue-jay/blueshift/domain"
	"github.com/blue-jay/blueshift/lib/flight"
	"github.com/blue-jay/blueshift/model/user"
	"github.com/gorilla/securecookie"

	"github.com/blue-jay/core/passhash"
)

// RegisterHandler represents the services required for this controller.
type RegisterHandler struct {
	UserService domain.UserCase
	ViewService domain.ViewCase
}

// Index displays the register page.
func (h *RegisterHandler) Index(w http.ResponseWriter, r *http.Request) {
	v := h.ViewService.New("register/index")
	//form.Repopulate(r.Form, v.Vars, "first_name", "last_name", "email")

	/*expiration := time.Now().Add(365 * 24 * time.Hour) ////Set to expire in 1 year
	cookie := http.Cookie{Name: "_blueshift", Value: "alice_cooper@gmail.com", Expires: expiration, HttpOnly: false}
	http.SetCookie(w, &cookie)*/
	saveSession(w, r)
	v.Render(w, r)
}

func saveSession(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	// Decode authentication key
	auth, err := base64.StdEncoding.DecodeString(c.Config.Session.AuthKey)
	if err != nil || len(auth) == 0 {
		log.Println(err)
	}
	encrypt, err := base64.StdEncoding.DecodeString(c.Config.Session.EncryptKey)
	if err != nil {
		log.Println(err)
	}

	var s = securecookie.New(auth, encrypt)

	var cookieName = "_blueshift"

	// Read the cookie.
	_, err = r.Cookie(cookieName)
	if err != nil {
		value := map[string]string{
			"foo": "bar",
		}

		if encoded, err := s.Encode(cookieName, value); err == nil {
			cookie := &http.Cookie{
				Name:     cookieName,
				Value:    encoded,
				Path:     "/",
				HttpOnly: false,
			}
			http.SetCookie(w, cookie)
			log.Println(cookie)
		} else {
			log.Println(err)
		}
	}
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
