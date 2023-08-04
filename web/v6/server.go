package v6

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// AddRoute 路由注册功能
	// method http方法
	// path 请求路径
	// handleFunc 业务逻辑
	addRoute(method string, path string, handleFunc HandleFunc)
}

var _ Server = &HTTPServer{}

type HTTPServer struct {
	router
	mdls []Middleware
}

type HTTPServerOption func(server *HTTPServer)

func NewHTTPServer(opts ...HTTPServerOption) *HTTPServer {
	res := &HTTPServer{
		router: newRouter(),
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func ServerWithMiddleware(mdls ...Middleware) HTTPServerOption {
	return func(server *HTTPServer) {
		server.mdls = mdls
	}
}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	// 查找路由并执行业务逻辑
	root := s.Serve
	for i := len(s.mdls) - 1; i >= 0; i-- {
		root = s.mdls[i](root)
	}
	// 第一个执行的 handle, 这里flashResp最后执行
	m := func(next HandleFunc) HandleFunc {
		return func(ctx *Context) {
			next(ctx)
			s.flashResp(ctx)
		}
	}
	root = m(root)
	root(ctx)
}

func (s *HTTPServer) Serve(ctx *Context) {
	mi, ok := s.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || mi.n == nil || mi.n.handler == nil {
		ctx.RespStatusCode = 404
		//ctx.Resp.WriteHeader(404)
		//ctx.Resp.Write([]byte("NOT FOUND"))
		return
	}
	ctx.PathParams = mi.pathParams
	ctx.MatchRoute = mi.n.route
	mi.n.handler(ctx)
}

func (s *HTTPServer) flashResp(ctx *Context) {
	if ctx.RespStatusCode > 0 {
		ctx.Resp.WriteHeader(ctx.RespStatusCode)
	}
	_, err := ctx.Resp.Write(ctx.RespData)
	if err != nil {
		log.Fatalln("回写响应失败", err)
	}
}

func (s *HTTPServer) Get(path string, handleFunc HandleFunc) {
	s.addRoute(http.MethodGet, path, handleFunc)
}

func (s *HTTPServer) Post(path string, handleFunc HandleFunc) {
	s.addRoute(http.MethodPost, path, handleFunc)
}

func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

/*func (s *HTTPServer) Start1(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 中间可以做一些业务的前置条件

	return http.Serve(listen, s)
}*/
