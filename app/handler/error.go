package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/root-san/root-san/pkg/dberror"
	"github.com/root-san/root-san/pkg/echoutil"
)

func catch(ec echo.Context, err error) error {
	if errors.Is(err, dberror.ErrDuplicateKey) {
		return ec.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg":          "already exist",
			"error_detail": err.Error(),
		})
	}
	if errors.Is(err, sql.ErrNoRows) {
		return ec.JSON(http.StatusNotFound, map[string]interface{}{
			"msg":          "not found",
			"error_detail": err.Error(),
		})
	}
	return echoutil.ErrInternal(ec, err)
}
