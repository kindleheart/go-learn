package cookie

import (
	"net/http"
)

type Propagator struct {
	cookieName   string
	cookieOption func(c *http.Cookie) // 用于对cookie设置
}

type PropagatorOption func(p *Propagator)

func NewPropagator(opts ...PropagatorOption) *Propagator {
	res := &Propagator{
		cookieName: "sessid",
		cookieOption: func(c *http.Cookie) {
		},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func WithCookieName(name string) PropagatorOption {
	return func(p *Propagator) {
		p.cookieName = name
	}
}

func WithCookieOption(cookieOption func(c *http.Cookie)) PropagatorOption {
	return func(p *Propagator) {
		p.cookieOption = cookieOption
	}
}

func (p *Propagator) Inject(id string, writer http.ResponseWriter) error {
	c := &http.Cookie{
		Name:  p.cookieName,
		Value: id,
	}
	p.cookieOption(c)
	http.SetCookie(writer, c)
	return nil
}

func (p *Propagator) Extract(req *http.Request) (string, error) {
	c, err := req.Cookie(p.cookieName)
	if err != nil {
		return "", err
	}
	return c.Value, nil
}

func (p *Propagator) Remove(writer http.ResponseWriter) error {
	c := &http.Cookie{
		Name:   p.cookieName,
		MaxAge: -1, // 代表直接删除cookie
	}
	http.SetCookie(writer, c)
	return nil
}
