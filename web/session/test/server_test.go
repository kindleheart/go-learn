package test

import (
	"goLearn/web/jie"
	"goLearn/web/jie/middleware/accesslog"
	"goLearn/web/jie/middleware/recover"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	user := struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Number int64  `json:"number"`
	}{
		Name: "jj",
		Age:  22,
	}
	accessLog := accesslog.NewBuilder().Build()
	recover := recover.NewMiddlewareBuilder().Build()
	r := jie.NewHTTPServer()
	r.Use(accessLog)
	r.Get("/a/b/c", func(ctx *jie.Context) {
		_ = ctx.RespString(http.StatusOK, "哈哈哈测试通过了")
	})
	r.Get("/a/*", func(ctx *jie.Context) {
		_ = ctx.RespString(http.StatusOK, "通配符匹配匹配成功")
	})
	g := r.Group("/user")
	g.Get("/info/:id", func(ctx *jie.Context) {
		user.Number, _ = ctx.PathValue("id").ToInt64()
		_ = ctx.RespJSON(http.StatusOK, user)
	})
	g.Get("/login", func(ctx *jie.Context) {
		panic("嘻嘻")
		_ = ctx.RespString(http.StatusOK, "登录成功")
	}, recover)
	_ = r.Start(":8080")
}
