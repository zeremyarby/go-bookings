package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware that adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfhandler := nosurf.New(next)
	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, // Set to true in production
		SameSite: http.SameSiteLaxMode,
	})

	return csrfhandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
