package accesslog

import (
	"encoding/json"
	"goLearn/web/jie"
	"log"
)

type MiddleWareBuilder struct {
	logFunc func(accessLog string)
}

func (b *MiddleWareBuilder) LogFunc(logFunc func(accessLog string)) *MiddleWareBuilder {
	b.logFunc = logFunc
	return b
}

func NewBuilder() *MiddleWareBuilder {
	return &MiddleWareBuilder{
		logFunc: func(accessLog string) {
			log.Println(accessLog)
		},
	}
}

type accessLog struct {
	Host       string `json:"host"`
	Route      string `json:"route"`
	HTTPMethod string `json:"http_method"`
	Path       string `json:"path"`
}

func (b *MiddleWareBuilder) Build() jie.Middleware {
	return func(next jie.HandleFunc) jie.HandleFunc {
		return func(ctx *jie.Context) {
			defer func() {
				l := accessLog{
					Host:       ctx.Req.URL.Host,
					Route:      ctx.MatchRoute,
					HTTPMethod: ctx.Req.Method,
					Path:       ctx.Req.URL.Path,
				}
				val, _ := json.Marshal(l)
				b.logFunc(string(val))
			}()
			next(ctx)
		}
	}
}
