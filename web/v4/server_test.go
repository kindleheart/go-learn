package v2

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	s.addRoute(http.MethodGet, "/user", func(ctx *Context) {
		_, _ = ctx.Resp.Write([]byte("hello world"))
	})
	s.addRoute(http.MethodGet, "/user/*", func(ctx *Context) {
		_, _ = ctx.Resp.Write([]byte("通配符匹配"))
	})
	s.addRoute(http.MethodGet, "/user/info/xixi", func(ctx *Context) {
		ctx.Resp.Write([]byte("我爱编程!!!"))
	})
	err := s.Start(":8080")
	if err != nil {
		return
	}
}
