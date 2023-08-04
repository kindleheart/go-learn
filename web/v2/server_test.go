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

	err := s.Start(":8080")
	if err != nil {
		return
	}
}
