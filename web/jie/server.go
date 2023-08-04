package jie

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)
type MiddlewaresChain []Middleware

var _ server = &HTTPServer{}

type server interface {
	http.Handler
	Start(addr string) error
	addRoute(method, path string, handler HandleFunc, middlewares MiddlewaresChain)
}

type HTTPServer struct {
	router
	RouterGroup
	tplEngine TemplateEngine
}

func NewHTTPServer(opts ...HTTPServerOption) *HTTPServer {
	server := &HTTPServer{
		router: newRouter(),
		RouterGroup: RouterGroup{
			basePath:    "",
			middlewares: nil,
		},
	}
	server.RouterGroup.server = server
	for _, opt := range opts {
		opt(server)
	}
	return server
}

type HTTPServerOption func(server *HTTPServer)

func ServerWithTemplateEngine(g *GoTemplateEngine) HTTPServerOption {
	return func(server *HTTPServer) {
		server.tplEngine = g
	}
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:       request,
		Resp:      writer,
		tplEngine: h.tplEngine,
	}
	nodeValue := h.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if nodeValue == nil {
		_ = ctx.RespString(http.StatusNotFound, "NOT FOUND")
		return
	}
	ctx.PathParams = nodeValue.pathParams
	ctx.MatchRoute = nodeValue.route

	// 以此执行中间件和handler
	root := nodeValue.handler
	mdls := nodeValue.middlewares
	for i := len(mdls) - 1; i >= 0; i-- {
		root = mdls[i](root)
	}

	// flashResp最后执行
	m := func(next HandleFunc) HandleFunc {
		return func(ctx *Context) {
			next(ctx)
			h.flashResp(ctx)
		}
	}
	root = m(root)
	root(ctx)
}

func (h *HTTPServer) flashResp(ctx *Context) {
	if ctx.RespStatusCode > 0 {
		ctx.Resp.WriteHeader(ctx.RespStatusCode)
	}
	if len(ctx.RespData) == 0 {
		return
	}
	_, err := ctx.Resp.Write(ctx.RespData)
	if err != nil {
		log.Fatalln("回写响应失败", err)
	}
}

func (h *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, h)
}
