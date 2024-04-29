package json

import (
	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/cat-social/shared/common"
)

type response struct {
}

func NewResponse() common.JSON {
	return &response{}
}

func (c *response) Ok(ctx echo.Context, message string, data interface{}) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
