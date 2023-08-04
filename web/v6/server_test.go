package v6

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	s.addRoute(http.MethodGet, "/user", func(ctx *Context) {
		ctx.RespJSON(http.StatusOK, &User{"xixi", 11})
	})
	s.addRoute(http.MethodGet, "/user/*", func(ctx *Context) {
		val := ctx.PathValue("id")
		if val.err != nil {
			ctx.Resp.Write([]byte("找不到路径参数"))
			return
		}
		_, _ = ctx.Resp.Write([]byte(fmt.Sprintf("路径参数: %s", val.val)))
	})
	s.addRoute(http.MethodGet, "/user/info/xixi", func(ctx *Context) {
		ctx.Resp.Write([]byte("我爱编程!!!"))
	})
	err := s.Start(":8080")
	if err != nil {
		return
	}
}
