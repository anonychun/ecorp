package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	c echo.Context

	body struct {
		Ok     bool `json:"ok"`
		Meta   any  `json:"meta"`
		Data   any  `json:"data"`
		Errors any  `json:"errors"`
	}
}

func NewResponse(c echo.Context) *Response {
	return &Response{c: c}
}

func (r *Response) SetStatus(status int) *Response {
	r.c.Response().Status = status
	return r
}

func (r *Response) SetMeta(meta any) *Response {
	r.body.Meta = meta
	return r
}

func (r *Response) SetData(data any) *Response {
	r.body.Data = data
	return r
}

func (r *Response) SetErrors(err error) *Response {
	switch e := err.(type) {
	case *Error:
		r.SetStatus(e.Status)
		message, ok := e.Errors.(string)
		if ok {
			r.body.Errors = echo.Map{"message": message}
		} else {
			r.body.Errors = e.Errors
		}
	case *echo.HTTPError:
		r.SetStatus(e.Code)
		r.body.Errors = echo.Map{"message": e.Message}
	default:
		r.SetStatus(http.StatusInternalServerError)
		r.body.Errors = echo.Map{"message": "Something went wrong"}
	}

	return r
}

func (r *Response) Send() error {
	r.body.Ok = r.c.Response().Status >= http.StatusOK && r.c.Response().Status < http.StatusMultipleChoices
	if r.body.Ok {
		r.body.Errors = nil
	} else {
		r.body.Data = nil
	}

	return r.c.JSON(r.c.Response().Status, r.body)
}

func (r *Response) SendMessage(message string) error {
	return r.SetData(echo.Map{"message": message}).Send()
}

func (r *Response) SendOk() error {
	return r.SetStatus(http.StatusOK).SendMessage("ok")
}
