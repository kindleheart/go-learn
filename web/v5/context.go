package v2

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter
	// 路由参数
	PathParams map[string]string
	// 查询参数缓存
	queryValues url.Values
}

// BindJSON body输入
func (c *Context) BindJSON(val any) error {
	if c.Req.Body == nil {
		return errors.New("web: body is nil")
	}
	// 不用unmarshal，.cReq.Body是一个stream
	decoder := json.NewDecoder(c.Req.Body)
	return decoder.Decode(val)
}

// FormValue 处理表单输入, 有缓存
func (c *Context) FormValue(key string) StringValue {
	err := c.Req.ParseForm()
	if err != nil {
		return StringValue{err: err}
	}
	return StringValue{val: c.Req.FormValue(key)}
}

// QueryValue 查询参数, Query和表单比起来，它没有缓存
func (c *Context) QueryValue(key string) StringValue {
	if c.queryValues == nil {
		c.queryValues = c.Req.URL.Query()
	}
	vals, ok := c.queryValues[key]
	if !ok {
		return StringValue{err: errors.New("web: key不存在")}
	}
	return StringValue{val: vals[0]}
}

// PathValue 路径参数
func (c *Context) PathValue(key string) StringValue {
	val, ok := c.PathParams[key]
	if !ok {
		return StringValue{err: errors.New("web: 找不到这个key")}
	}
	return StringValue{val: val}
}

func (c *Context) RespJSON(code int, val any) error {
	bs, err := json.Marshal(val)
	if err != nil {
		return err
	}
	c.Resp.WriteHeader(code)
	c.Resp.Header().Set("Content-Type", "application/json")
	_, err = c.Resp.Write(bs)
	return err

}

func (c *Context) RespJSONOK(val any) error {
	return c.RespJSON(http.StatusOK, val)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Resp, cookie)
}

type StringValue struct {
	val string
	err error
}

func (s StringValue) ToInt64() (int64, error) {
	if s.err != nil {
		return 0, s.err
	}
	return strconv.ParseInt(s.val, 10, 64)
}
