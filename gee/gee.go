package gee

import "net/http"

// HandlerFunc defines the handler used by gee
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engin *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engin.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {

}
