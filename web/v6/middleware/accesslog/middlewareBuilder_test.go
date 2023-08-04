package accesslog

import (
	"fmt"
	v6 "goLearn/web/v6"
	"net/http"
	"testing"
)

func TestMiddlewareBuilder(t *testing.T) {
	builder := MiddleWareBuilder{}
	mdl := builder.LogFunc(func(accessLog string) {
		fmt.Println(accessLog)
	}).build()

	server := v6.NewHTTPServer(v6.ServerWithMiddleware(mdl))
	server.Get("/a/b/*", func(ctx *v6.Context) {
		fmt.Println("hello, it is me")
	})
	req, err := http.NewRequest(http.MethodGet, "/a/b/c", nil)
	if err != nil {
		t.Fatal(err)
	}
	server.ServeHTTP(nil, req)
}
