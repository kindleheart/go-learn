package jie

import (
	"html/template"
	"testing"
)

/*
	func TestHTTPServer_ServeHTTP(t *testing.T) {
		user := struct {
			Name string `json:"name,omitempty"`
			Age  int    `json:"age,omitempty"`
		}{
			Name: "jj",
			Age:  22,
		}
		accessLog := accesslog.NewBuilder().Build()
		r := NewHTTPServer()
		r.Get("/index", func(ctx *Context) {
			_ = ctx.RespString(http.StatusOK, "哈哈哈测试通过了")
		})

		g := r.Group("/user", accessLog)
		g.Get("/info", func(ctx *Context) {
			_ = ctx.RespJSON(http.StatusOK, user)
		})
		g.Get("/login", func(ctx *Context) {
			_ = ctx.RespString(http.StatusOK, "登录成功")
		})
		r.Start(":8080")
	}
*/
func TestServerWithRenderEngine(t *testing.T) {
	user := struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}{
		Name: "xixi",
		Age:  22,
	}
	tpl, err := template.ParseGlob("testdata/tpls/*.gohtml")
	if err != nil {
		t.Fatal(err)
	}
	s := NewHTTPServer(ServerWithTemplateEngine(&GoTemplateEngine{T: tpl}))
	s.Get("/login", func(ctx *Context) {
		er := ctx.Render("login.gohtml", user)
		if er != nil {
			t.Fatal(er)
		}
	})
	err = s.Start(":8081")
	if err != nil {
		t.Fatal(err)
	}
}
