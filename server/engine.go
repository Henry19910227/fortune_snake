package server

import (
	"context"
	"game_server_slots_fortune_snake/model"
)

type ResolveFunc func(b []byte) string

type Engine struct {
	RouterGroup
	route      *Router
	resolver   ResolveFunc
	config     model.ServerConfig
	grpcServer *Server
}

func New(config model.ServerConfig) *Engine {
	e := &Engine{
		RouterGroup: RouterGroup{
			Handlers: make([]HandlerFunc, 0),
			basePath: "",
		},
		route:  NewRouter(),
		config: config,
	}
	grpcServer := newServer(config, e)
	e.grpcServer = grpcServer
	e.RouterGroup.engine = e
	return e
}

func (e *Engine) Use(middleware HandlerFunc) {
	e.RouterGroup.Use(middleware)
}

func (e *Engine) addRoute(path string, handlers []HandlerFunc) {
	e.route.Add(path, handlers)
}

func (e *Engine) getRoute(path string) HandlerFunc {
	return nil
}

func (e *Engine) PathResolver(resolver ResolveFunc) {
	e.resolver = resolver
}

func (e *Engine) Run(ctx context.Context) {
	e.grpcServer.RunWithRetry(ctx)
}
