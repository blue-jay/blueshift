package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blue-jay/blueshift/app"
	"github.com/blue-jay/blueshift/app/mock"
)

func TestHandler(t *testing.T) {
	// Inject our mock into our handler.
	var us mock.UserService
	var vs mock.ViewService
	var h Handler
	h.UserService = &us
	h.ViewService = &vs

	// Mock our User() call.
	us.UserByEmailFn = func(email string) (*app.User, error) {
		if email != "example@example.com" {
			t.Fatalf("unexpected email: %v", email)
		}
		return &app.User{Email: "example@example.com", FirstName: "joe"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/test", nil)
	h.ServeHTTP(w, r)

	// Validate mock.
	if !us.UserByEmailInvoked {
		t.Fatal("expected UserByEmail() to be invoked")
	}
}
