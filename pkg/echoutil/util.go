package echoutil

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errOption struct {
	msg map[string]interface{}
}

type ErrOptionFunc func(*errOption)

func ErrDetail(msg map[string]interface{}) ErrOptionFunc {
	return func(eo *errOption) {
		eo.msg = msg
	}
}

func ErrInternal(ec echo.Context, err error, opts ...ErrOptionFunc) error {
	opt := errOption{}
	for _, f := range opts {
		f(&opt)
	}
	return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error":        "internal server error",
		"error_detail": err.Error(),
		"msg":          opt.msg,
	})
}

func ErrBadRequest(ec echo.Context, err error, opts ...ErrOptionFunc) error {
	opt := errOption{}
	for _, f := range opts {
		f(&opt)
	}
	return ec.JSON(http.StatusBadRequest, map[string]interface{}{
		"error":        "invalid request",
		"error_detail": err.Error(),
		"msg":          opt.msg,
	})
}
