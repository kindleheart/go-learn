package errorhdl

import (
	v6 "goLearn/web/v6"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	buildr := NewMiddlewareBuilder()
	m := buildr.RegisterError(404, []byte(`
<html>
	<h1>404 NOT FOUND</h1>
</html>
`)).Build()
	server := v6.NewHTTPServer(v6.ServerWithMiddleware(m))
	server.Get("/a/b/c", func(ctx *v6.Context) {
		ctx.Resp.Write([]byte("嘻嘻嘻"))
	})
	server.Start(":8080")
}
