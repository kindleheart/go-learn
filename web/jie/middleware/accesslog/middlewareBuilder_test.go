package accesslog

import (
	"fmt"
	"goLearn/web/jie"
	"net/http"
	"testing"
)

func TestMiddlewareBuilder(t *testing.T) {
	builder := MiddleWareBuilder{}
	mdl := builder.LogFunc(func(accessLog string) {
		fmt.Println(accessLog)
	}).Build()

	r := jie.NewHTTPServer()
	r.Get("/a/b/*", func(ctx *jie.Context) {
		fmt.Println("hello, it is me")
	}, mdl)
	req, err := http.NewRequest(http.MethodGet, "/a/b/c", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(nil, req)
}
