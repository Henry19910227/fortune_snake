package server

type ResolveFunc func(b []byte) string

type Engine struct {
	RouterGroup
	route    *Router
	resolver ResolveFunc
}

func New() *Engine {
	e := &Engine{
		RouterGroup: RouterGroup{
			Handlers: make([]HandlerFunc, 0),
			basePath: "",
		},
		route: NewRouter(),
	}
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
