package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/cat-social/shared/common"
	jwtAccess "github.com/mrakhaf/cat-social/shared/common/jwt"
)

type handlerCat struct {
	Json      common.JSON
	JwtAccess *jwtAccess.JWT
}

func CatHandler(catRoute *echo.Group, Json common.JSON, JwtAccess *jwtAccess.JWT) {
	handler := handlerCat{
		Json:      Json,
		JwtAccess: JwtAccess,
	}

	catRoute.GET("/cat", handler.GetCat)
}

func (h handlerCat) GetCat(c echo.Context) error {

	userId, err := h.JwtAccess.GetUserIdFromToken(c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome " + userId})
}
