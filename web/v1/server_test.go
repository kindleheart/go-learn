package v1

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var s *HTTPServer
	s.AddRoute(http.MethodGet, "/user", func(ctx *Context) {
		fmt.Fprint(ctx.Resp, "Hello World")
	})
	err := s.Start(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
