package jie

import (
	"net/http"
)

type RouterGroup struct {
	server      *HTTPServer
	basePath    string
	middlewares MiddlewaresChain
}

func (g *RouterGroup) Group(relativePath string, middlewares ...Middleware) *RouterGroup {
	return &RouterGroup{
		server:      g.server,
		basePath:    g.calculateAbsolutePath(relativePath),
		middlewares: g.combineMiddlewares(middlewares),
	}
}

func (g *RouterGroup) Use(middlewares ...Middleware) *RouterGroup {
	g.middlewares = append(g.middlewares, middlewares...)
	return g
}

func (g *RouterGroup) handle(method, relativePath string, handler HandleFunc, middlewares MiddlewaresChain) {
	absolutePath := g.calculateAbsolutePath(relativePath)
	middlewares = g.combineMiddlewares(middlewares)
	g.server.addRoute(method, absolutePath, handler, middlewares)
}

func (g *RouterGroup) Get(path string, handler HandleFunc, middlewares ...Middleware) {
	g.handle(http.MethodGet, path, handler, middlewares)
}

func (g *RouterGroup) Post(path string, handler HandleFunc, middlewares ...Middleware) {
	g.handle(http.MethodPost, path, handler, middlewares)
}

func (g *RouterGroup) combineMiddlewares(middlewares MiddlewaresChain) MiddlewaresChain {
	finalSize := len(g.middlewares) + len(middlewares)
	if middlewaresMaxNums < finalSize {
		panic("too many middlewares")
	}
	mergedMiddlewares := make(MiddlewaresChain, finalSize)
	copy(mergedMiddlewares, g.middlewares)
	copy(mergedMiddlewares[len(g.middlewares):], middlewares)
	return mergedMiddlewares
}

func (g *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return g.basePath + relativePath
}
