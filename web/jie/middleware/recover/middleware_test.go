package recover

import (
	"goLearn/web/jie"
	"testing"
)

func TestNewMiddlewareBuilder_Build(t *testing.T) {
	builder := NewMiddlewareBuilder()
	server := jie.NewHTTPServer()
	server.Get("/index", func(ctx *jie.Context) {
		panic("我panic了")
	}, builder.build())
	server.Start(":8080")
}
