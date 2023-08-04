package recover

import (
	v6 "goLearn/web/v6"
	"testing"
)

func TestNewMiddlewareBuilder_Build(t *testing.T) {
	builder := NewMiddlewareBuilder()
	server := v6.NewHTTPServer(v6.ServerWithMiddleware(builder.build()))
	server.Get("/index", func(ctx *v6.Context) {
		panic("我panic了")
	})
	server.Start(":8080")
}
