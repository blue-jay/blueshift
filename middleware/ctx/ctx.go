package ctx

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/blue-jay/blueshift/lib/flight"
	"github.com/gorilla/securecookie"
)

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		//var hashKey = []byte("very-secret")
		//var blockKey = []byte("a-lot-secret")

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
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		value := make(map[string]string)
		if err = s.Decode(cookieName, cookie.Value, &value); err == nil {
			//fmt.Fprintf(w, "The value of foo is %q", value["foo"])
			log.Printf("The value of foo is %q", value["foo"])
		} else {
			log.Println(value)
		}

		//Add data to context
		ctx := context.WithValue(r.Context(), "id", cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
