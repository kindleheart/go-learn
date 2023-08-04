package test

import (
	"fmt"
	"goLearn/web/jie"
	"goLearn/web/session"
	"goLearn/web/session/cookie"
	"goLearn/web/session/memory"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestSession(t *testing.T) {
	m := session.Manager{
		Propagator: cookie.NewPropagator(),
		Store:      memory.NewStore(time.Minute * 15),
		CtxSessKey: "sesskey",
	}
	r := jie.NewHTTPServer()
	// session中间件
	r.Use(func(next jie.HandleFunc) jie.HandleFunc {
		return func(ctx *jie.Context) {
			fmt.Println(ctx.Req.URL.Path)
			if ctx.Req.URL.Path == "/login" {
				next(ctx)
				return
			}
			_, err := m.GetSession(ctx)
			if err != nil {
				ctx.RespStatusCode = http.StatusUnauthorized
				ctx.RespData = []byte("请重新登录")
				return
			}
			// 刷新Session的过期时间
			err = m.RefreshSession(ctx)
			if err != nil {
				log.Println("刷新Session时间失败")
			}
			next(ctx)
		}
	})
	r.Post("/login", func(ctx *jie.Context) {
		// 在这之前校验用户名和密码
		sess, err := m.InitSession(ctx)
		if err != nil {
			ctx.RespStatusCode = http.StatusUnauthorized
			ctx.RespData = []byte("登录失败了")
			return
		}
		err = sess.Set(ctx.Req.Context(), "nickname", "jay")
		if err != nil {
			ctx.RespStatusCode = http.StatusUnauthorized
			ctx.RespData = []byte("登录失败了")
			return
		}
		ctx.RespStatusCode = http.StatusOK
		ctx.RespData = []byte("登录成功")
	})
	r.Post("/logout", func(ctx *jie.Context) {
		err := m.RemoveSession(ctx)
		if err != nil {
			ctx.RespStatusCode = http.StatusUnauthorized
			ctx.RespData = []byte("退出失败了")
			return
		}
		ctx.RespStatusCode = http.StatusOK
		ctx.RespData = []byte("退出登录成功")
	})
	r.Get("/user", func(ctx *jie.Context) {
		sess, _ := m.GetSession(ctx)
		val, _ := sess.Get(ctx.Req.Context(), "nickname")
		ctx.RespData = []byte(val.(string))
	})
	r.Start(":8080")
}
