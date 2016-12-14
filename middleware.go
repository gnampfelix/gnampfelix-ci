package main

import (
	"net/http"
)

//	A Middleware combines multiple Middleware
type Middleware []http.Handler
type MiddlewareResponseWriter struct {
	http.ResponseWriter
	isWritten bool
}

func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}

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
