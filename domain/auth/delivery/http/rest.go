package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
	"github.com/mrakhaf/cat-social/models/request"
)

type handlerAuth struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
}

func AuthHandler(authRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository) {
	handler := handlerAuth{usecase: usecase, repository: repository}

	authRoute.POST("/login", handler.Login)
}

func (h *handlerAuth) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := h.usecase.Login(c.Request().Context(), req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, token)

}
