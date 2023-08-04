package errorhdl

import (
	"goLearn/web/jie"
	"net/http"
	"testing"
)

func TestMiddlewareBuilder_Build(t *testing.T) {
	buildr := NewMiddlewareBuilder()
	m := buildr.RegisterError(http.StatusBadRequest, []byte(`
<html>
	<h1>Bad Request</h1>
</html>
`)).Build()
	server := jie.NewHTTPServer()
	server.Get("/a/b/c", func(ctx *jie.Context) {
		ctx.RespStatusCode = http.StatusBadRequest
	}, m)
	server.Start(":8080")
}
