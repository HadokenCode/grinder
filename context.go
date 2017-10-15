package grinder

import (
	"encoding/json"
	"net/http"
)

type (
	// Context interface
	Context interface {
		Request() *http.Request
		Response() *Response
		JSON(int, interface{}) error
		String(int, string) error
		Code(int) error
		HTTPError(int, string) error
		AddParams(map[string]string)
		GetParam(string) string
		GetParams() map[string]string
		HasParam(string) bool
		SetHeader(string, string)
		GetHeader(string) string
		Redirect(int, string) error
	}

	context struct {
		request  *http.Request
		response *Response
		params   map[string]string
	}
)

func (c *context) Request() *http.Request {
	return c.request
}

func (c *context) Response() *Response {
	return c.response
}

func (c *context) JSON(code int, i interface{}) (err error) {
	b, err := json.Marshal(i)

	if err != nil {
		c.HTTPError(500, err.Error())
		return
	}

	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(code)
	_, err = c.response.Write([]byte(b))
	return
}

func (c *context) String(code int, s string) (err error) {
	c.response.Header().Set("Content-Type", "text/html;charset=utf-8")
	c.response.WriteHeader(code)
	_, err = c.response.Write([]byte(s))
	return
}

func (c *context) Code(code int) (err error) {
	c.response.WriteHeader(code)
	return nil
}

func (c *context) HTTPError(code int, message string) (err error) {
	c.response.Header().Set("Content-Type", "text/html;charset=utf-8")
	c.response.WriteHeader(code)
	_, err = c.response.Write([]byte(message))
	return
}

func (c *context) Redirect(code int, uri string) (err error) {
	http.Redirect(c.Response().writer, c.Request(), uri, code)
	return nil
}

func (c *context) AddParams(params map[string]string) {
	if c.params == nil {
		c.params = make(map[string]string)
	}

	for k, v := range params {
		c.params[k] = v
	}

	return
}

func (c *context) GetParam(i string) string {
	param := c.params[i]
	return param
}

func (c *context) GetParams() map[string]string {
	return c.params
}

func (c *context) HasParam(i string) bool {
	_, isset := c.params[i]
	return isset
}

func (c *context) SetHeader(k string, v string) {
	c.response.Header().Set(k, v)
}

func (c *context) GetHeader(k string) string {
	return c.request.Header.Get(k)
}
