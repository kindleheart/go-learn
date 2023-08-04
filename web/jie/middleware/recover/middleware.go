package recover

import (
	"fmt"
	"goLearn/web/jie"
)

type MiddlewareBuilder struct {
	StatusCode int
	ErrMsg     string
	LogFunc    func(ctx *jie.Context)
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		StatusCode: 500,
		ErrMsg:     "you are panic",
		LogFunc: func(ctx *jie.Context) {
			fmt.Printf("panic路径：%s\n", ctx.Req.URL.String())
		},
	}
}

func (m *MiddlewareBuilder) Build() jie.Middleware {
	return func(next jie.HandleFunc) jie.HandleFunc {
		return func(ctx *jie.Context) {
			defer func() {
				if err := recover(); err != nil {
					ctx.RespStatusCode = m.StatusCode
					ctx.RespData = []byte(m.ErrMsg)
					m.LogFunc(ctx)
				}
			}()
			next(ctx)
		}
	}
}
