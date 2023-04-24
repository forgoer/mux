package mux

import (
	"net/http"
)

// Mux is an HTTP request multiplexer.
type Mux struct {
	http.ServeMux
	beforeHandlers []http.Handler
	afterHandlers  []http.Handler
	recoverHandler func(w http.ResponseWriter, r *http.Request)
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mux.recoverHandler != nil {
		defer mux.recoverHandler(w, r)
	}
	mux.ServeMux.ServeHTTP(w, r)
}

// HandleFunc registers the handler function for the given pattern.
func (mux *Mux) HandleFunc(pattern string, handlerFuns ...func(http.ResponseWriter, *http.Request)) {
	if len(handlerFuns) == 0 {
		panic("http: empty handlers")
	}
	handlers := make([]http.Handler, len(handlerFuns))
	for i, hf := range handlerFuns {
		handlers[i] = http.HandlerFunc(hf)
	}
	mux.Handle(pattern, handlers...)
}

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *Mux) Handle(pattern string, handlers ...http.Handler) {
	if len(handlers) == 0 {
		panic("http: empty handlers")
	}
	handlers = append(mux.beforeHandlers, handlers...)
	handlers = append(handlers, mux.afterHandlers...)
	mux.ServeMux.Handle(pattern, Handlers(handlers))
}

// Before set the handler before all handlers.
func (mux *Mux) Before(handler func(w http.ResponseWriter, r *http.Request)) {
	mux.beforeHandlers = append(mux.beforeHandlers, http.HandlerFunc(handler))
}

// After set the handler after all handlers.
func (mux *Mux) After(handler func(w http.ResponseWriter, r *http.Request)) {
	mux.afterHandlers = append(mux.afterHandlers, http.HandlerFunc(handler))
}

// Recover set the recover handler
func (mux *Mux) Recover(handler func(w http.ResponseWriter, r *http.Request)) {
	mux.recoverHandler = handler
}
