package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range append(m.globalMiddlewares, middlewares...) {
		n = middleware(n)
	}
	return n
}
func (m *Manager) WrapMux(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range append(m.globalMiddlewares, middlewares...) {
		n = middleware(n)
	}
	return n
}
