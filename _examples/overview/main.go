package main

import (
	"fmt"
	"net/http"

	"github.com/skynet0590/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(w http.ResponseWriter, r *http.Request) {

	// Check if user is authenticated
	if auth, _ := sess.Start(w, r).GetBoolean("authenticated"); !auth {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Print secret message
	w.Write([]byte("The cake is a lie!"))
}

func login(w http.ResponseWriter, r *http.Request) {
	session := sess.Start(w, r)

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Set("authenticated", true)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session := sess.Start(w, r)

	// Revoke users authentication
	session.Set("authenticated", false)
}

func scan(w http.ResponseWriter, r *http.Request) {
	c := 0
	sess.Scan(func(s *sessions.Session) {
		c++
		fmt.Println(s.Get("authenticated"))
	})
	w.Write([]byte(fmt.Sprintf("The sessions length is %d", c)))
}

func main() {
	app := http.NewServeMux()
	app.HandleFunc("/scan", scan)
	app.HandleFunc("/secret", secret)
	app.HandleFunc("/login", login)
	app.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", app)
}
