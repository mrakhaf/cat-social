package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/cat-social/domain/match_cat/interfaces"
	"github.com/mrakhaf/cat-social/models/request"
	"github.com/mrakhaf/cat-social/shared/common"
	jwtAccess "github.com/mrakhaf/cat-social/shared/common/jwt"
)

type handlerMatchCat struct {
	JSON      common.JSON
	JwtAccess *jwtAccess.JWT
	usecase   interfaces.Usecase
}

func MatchCatHandler(matchCatRoute *echo.Group, JSON common.JSON, JwtAccess *jwtAccess.JWT, usecase interfaces.Usecase) {
	handler := handlerMatchCat{
		JSON:      JSON,
		JwtAccess: JwtAccess,
		usecase:   usecase,
	}

	matchCatRoute.POST("/match																																																																																", handler.MatchCat)
}

func (h handlerMatchCat) MatchCat(c echo.Context) error {

	var matchCat request.MatchCat

	if err := c.Bind(&matchCat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(matchCat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})

	}

	userId, err := h.JwtAccess.GetUserIdFromToken(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	data, err := h.usecase.SaveMatchCat(c.Request().Context(), matchCat, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.JSON.Ok(c, "success", data)
}
