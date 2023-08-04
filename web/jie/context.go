package jie

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const middlewaresMaxNums int = math.MaxInt8 >> 1

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter

	// 路由参数
	PathParams map[string]string

	//  查询参数
	QueryValues url.Values

	// 匹配的路由
	MatchRoute string

	// 用于错误处理中间件
	RespStatusCode int
	RespData       []byte

	// 模板引擎
	tplEngine TemplateEngine

	// session数据缓存
	UserValues map[string]any
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
		return StringValue{Err: err}
	}
	return StringValue{Val: c.Req.FormValue(key)}
}

// QueryValue 查询参数, Query和表单比起来，它没有缓存
func (c *Context) QueryValue(key string) StringValue {
	if c.QueryValues == nil {
		c.QueryValues = c.Req.URL.Query()
	}
	vals, ok := c.QueryValues[key]
	if !ok {
		return StringValue{Err: errors.New("web: key不存在")}
	}
	return StringValue{Val: vals[0]}
}

// PathValue 路径参数
func (c *Context) PathValue(key string) StringValue {
	val, ok := c.PathParams[key]
	if !ok {
		return StringValue{Err: errors.New("web: 找不到这个key")}
	}
	return StringValue{Val: val}
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

func (c *Context) RespString(code int, val string) error {
	c.Resp.WriteHeader(code)
	_, err := c.Resp.Write([]byte(val))
	return err
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Resp, cookie)
}

func (c *Context) Render(tplName string, data any) error {
	var err error
	c.RespData, err = c.tplEngine.Render(c.Req.Context(), tplName, data)
	c.RespStatusCode = http.StatusOK
	if err != nil {
		c.RespStatusCode = 500
	}
	return nil
}

type StringValue struct {
	Val string
	Err error
}

func (s StringValue) String() (string, error) {
	if s.Err != nil {
		return "", s.Err
	}
	return s.Val, nil
}

func (s StringValue) ToInt64() (int64, error) {
	if s.Err != nil {
		return 0, s.Err
	}
	return strconv.ParseInt(s.Val, 10, 64)
}
