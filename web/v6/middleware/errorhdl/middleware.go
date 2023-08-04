package errorhdl

import v6 "goLearn/web/v6"

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

func (m *MiddlewareBuilder) Build() v6.Middleware {
	return func(next v6.HandleFunc) v6.HandleFunc {
		return func(ctx *v6.Context) {
			next(ctx)
			resp, ok := m.resp[ctx.RespStatusCode]
			if ok {
				ctx.RespData = resp
			}
		}
	}
}
