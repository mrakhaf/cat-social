package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/cat-social/domain/auth/interfaces"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/shared/common"
)

type handlerAuth struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
}

func AuthHandler(authRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, Json common.JSON) {
	handler := handlerAuth{
		usecase:    usecase,
		repository: repository,
		Json:       Json,
	}

	authRoute.POST("/login", handler.Login)

	authRoute.GET("/test", handler.Test)
}

func (h *handlerAuth) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		fmt.Println("test1")
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(&req); err != nil {
		fmt.Println("test2")
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err := h.usecase.Login(c.Request().Context(), req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return h.Json.Ok(c, "User logged successfully", data)

}

func (h *handlerAuth) Test(c echo.Context) error {

	return c.String(http.StatusOK, "Welcome to Cat Social!!!")

}
