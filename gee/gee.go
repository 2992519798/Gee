package gee

import (
	"net/http"
)

// HandlerFunc defines the handler used by gee
// 提供给框架用户，用来定义路由映射的处理方法
type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engin := &Engine{router: newRouter()}
	engin.groups = []*RouterGroup{engin.RouterGroup}
	engin.RouterGroup = &RouterGroup{engine: engin}
	return engin
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRouter("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRouter("POST", pattern, handler)
}

// RUN  warps the method to start a http server
func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP implements the interface of a  Handler (Engine)
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
