package v1

import (
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
	AddRoute(method string, path string, handleFunc HandleFunc)
}

var _ Server = &HTTPServer{}

type HTTPServer struct {
}

func (s *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
}

func (s *HTTPServer) Get(path string, handleFunc HandleFunc) {
	s.AddRoute(http.MethodGet, path, handleFunc)
}

func (s *HTTPServer) Post(path string, handleFunc HandleFunc) {
	s.AddRoute(http.MethodPost, path, handleFunc)
}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	// 查找路由并执行业务逻辑
	s.Serve(ctx)
}

func (s *HTTPServer) Serve(ctx *Context) {
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
