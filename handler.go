package mux

import (
	"context"
	"net/http"
)

var stoppedKey = "http-context-stopped"

// Handlers A collection of http.Handler responds to an HTTP request.
type Handlers []http.Handler

// ServeHTTP calls f(w, r).
func (hs Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, h := range hs {
		if Stopped(r) {
			break
		}
		h.ServeHTTP(w, r)
	}
}

func Stop(r *http.Request) {
	if r == nil {
		return
	}
	r1 := r.Clone(context.WithValue(r.Context(), &stoppedKey, true))
	*r = *r1
}

func Stopped(r *http.Request) bool {
	if r == nil {
		return false
	}
	if stopped, ok := r.Context().Value(&stoppedKey).(bool); ok && stopped {
		return true
	}

	return false
}
