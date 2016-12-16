//	Package to handle common routing operations.
package middleware

import (
	"net/http"
)

//	A Middleware combines multiple http.Handler
type Middleware []http.Handler
type MiddlewareResponseWriter struct {
	http.ResponseWriter
	isWritten bool
}

//	Create a new, empty Middleware.
func New() Middleware {
	return Middleware{}
}

//	Add an existing http.Handler to the middleware.
func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}

//	Iterates over each HTTP-Handler. If a handler reacts and writes to the response,
//	the loop stops. If no handler reacts, http.NotFound() is called.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw := NewMiddlewareResponseWriter(w)

	for _, handler := range m {
		handler.ServeHTTP(mw, r)
		if mw.isWritten {
			return
		}
	}
	http.NotFound(w, r)
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter{
		ResponseWriter: w,
	}
}

func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.isWritten = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.isWritten = true
	w.ResponseWriter.WriteHeader(code)
}
