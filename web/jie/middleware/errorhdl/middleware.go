package errorhdl

import (
	"goLearn/web/jie"
)

type MiddlewareBuilder struct {
	resp map[int][]byte
}

func NewMiddlewareBuilder() *MiddlewareBuilder {
	return &MiddlewareBuilder{resp: make(map[int][]byte, 64)}
}

func (m *MiddlewareBuilder) RegisterError(code int, resp []byte) *MiddlewareBuilder {
	m.resp[code] = resp
	return m
}

func (m *MiddlewareBuilder) Build() jie.Middleware {
	return func(next jie.HandleFunc) jie.HandleFunc {
		return func(ctx *jie.Context) {
			next(ctx)
			resp, ok := m.resp[ctx.RespStatusCode]
			if ok {
				ctx.RespData = resp
			}
		}
	}
}
