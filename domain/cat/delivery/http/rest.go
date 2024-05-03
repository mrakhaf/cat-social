package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/cat-social/domain/cat/interfaces"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/shared/common"
	jwtAccess "github.com/mrakhaf/cat-social/shared/common/jwt"
	"github.com/mrakhaf/cat-social/shared/utils"
)

type handlerCat struct {
	Json      common.JSON
	JwtAccess *jwtAccess.JWT
	usecase   interfaces.Usecase
}

func CatHandler(catRoute *echo.Group, Json common.JSON, JwtAccess *jwtAccess.JWT, usecase interfaces.Usecase) {
	handler := handlerCat{
		Json:      Json,
		JwtAccess: JwtAccess,
		usecase:   usecase,
	}

	catRoute.POST("", handler.UploadCat)
	catRoute.GET("", handler.GetCats)
	catRoute.PUT("/:id", handler.UpdateCat)
	catRoute.DELETE("/:id", handler.DeleteCat)
}

func (h handlerCat) UploadCat(c echo.Context) error {

	userId, err := h.JwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	var dataUploadCat request.UploadCat

	if err := c.Bind(&dataUploadCat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(dataUploadCat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	isRaceTrue := utils.ValidationRace(dataUploadCat.Race)

	if !isRaceTrue {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Wrong race"})
	}

	data, err := h.usecase.UploadCat(c.Request().Context(), dataUploadCat, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.Ok(c, "success", data)
}

func (h handlerCat) GetCats(c echo.Context) error {

	userId, err := h.JwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	var req request.GetCatParam

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return err
	}

	fmt.Println(req)

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if req.Race != nil {

		isRace := utils.ValidationRace(*req.Race)

		if !isRace {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Wrong race"})
		}
	}

	data, err := h.usecase.GetCat(c.Request().Context(), userId, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.Ok(c, "success", data)

}

func (h handlerCat) UpdateCat(c echo.Context) error {

	var req request.UploadCat

	idCat := c.Param("id")

	fmt.Println(idCat)

	userId, err := h.JwtAccess.GetUserIdFromToken(c)

	fmt.Println(userId)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	isRaceTrue := utils.ValidationRace(req.Race)

	if !isRaceTrue {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Wrong race"})
	}

	if err := h.usecase.ValidationCatUser(c.Request().Context(), idCat, userId); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "id is not found"})
	}

	err = h.usecase.UpdateCat(c.Request().Context(), idCat, userId, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.Ok(c, "success successfully add cat", nil)
}

func (h handlerCat) DeleteCat(c echo.Context) error {
	id := c.Param("id")

	userId, err := h.JwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	if err := h.usecase.ValidationCatUser(c.Request().Context(), id, userId); err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusNotFound, map[string]string{"message": "id is not found"})
	}

	err = h.usecase.DeleteCat(c.Request().Context(), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.Ok(c, "success successfully delete cat", nil)
}
