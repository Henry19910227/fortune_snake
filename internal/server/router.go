package server

import "sync"

type HandlerFunc func(c *Context)

type Router struct {
	route map[string][]HandlerFunc
	mu    sync.RWMutex
}

func NewRouter() *Router {
	return &Router{
		route: make(map[string][]HandlerFunc),
	}
}

func (r *Router) Add(path string, handlers []HandlerFunc) {
	r.mu.Lock()
	defer func() {
		r.mu.Unlock()
	}()
	r.route[path] = handlers
}

func (r *Router) Get(path string) ([]HandlerFunc, bool) {
	r.mu.RLock()
	defer func() {
		r.mu.RUnlock()
	}()
	handler, ok := r.route[path]
	if !ok {
		return nil, false
	}
	return handler, true
}
