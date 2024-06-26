package http

import (
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
	authRoute.POST("/register", handler.Register)
}

func (h *handlerAuth) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//check is email exist
	isEmailExist, dataUser, err := h.usecase.CheckIsEmailExist(c.Request().Context(), req.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if !isEmailExist {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Email not found"})
	}

	//check password
	isPasswordCorrect := h.usecase.CheckUserPassword(c.Request().Context(), req.Password, dataUser)

	if !isPasswordCorrect {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wrong password"})
	}

	data, err := h.usecase.Login(c.Request().Context(), req.Email, req.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.Ok(c, "User logged successfully", data)

}

func (h *handlerAuth) Register(c echo.Context) error {

	var req request.Register

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//validate email
	isEmailExist, _, err := h.usecase.CheckIsEmailExist(c.Request().Context(), req.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if isEmailExist {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email already exist"})
	}

	data, err := h.usecase.Register(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully",
		"data":    data,
	})
}
