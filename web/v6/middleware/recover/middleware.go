package recover

import (
	"fmt"
	v6 "goLearn/web/v6"
)

type MiddlewareBuilder struct {
	StatusCode int
	ErrMsg     string
	LogFunc    func(ctx *v6.Context)
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{
		StatusCode: 500,
		ErrMsg:     "you are panic",
		LogFunc: func(ctx *v6.Context) {
			fmt.Printf("panic路径：%s", ctx.Req.URL.String())
		},
	}
}

func (m *MiddlewareBuilder) build() v6.Middleware {
	return func(next v6.HandleFunc) v6.HandleFunc {
		return func(ctx *v6.Context) {
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
