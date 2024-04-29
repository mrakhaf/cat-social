package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authHandler "github.com/mrakhaf/cat-social/domain/auth/delivery/http"
	"github.com/mrakhaf/cat-social/domain/auth/repository"
	"github.com/mrakhaf/cat-social/domain/auth/usecase"
	"github.com/mrakhaf/cat-social/shared/common"
	formatJson "github.com/mrakhaf/cat-social/shared/common/json"
	"github.com/mrakhaf/cat-social/shared/common/jwt"
	"github.com/mrakhaf/cat-social/shared/config/database"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = common.NewValidator()

	//db config
	catDB, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Cat Social!!!")
	})

	//create group
	userGroup := e.Group("/v1/user")

	// catGroup := e.Group("/v1/cat")
	// {
	// 	config := echojwt.Config{
	// 		SigningKey: []byte("secret"),
	// 	}

	// 	catGroup.Use(echojwt.WithConfig(config))
	// }

	//
	formatResponse := formatJson.NewResponse()
	jwtAccess := jwt.NewJWT()

	//auth
	authRepo := repository.NewRepository(catDB)
	authUsecase := usecase.NewUsecase(authRepo, jwtAccess)
	authHandler.AuthHandler(userGroup, authUsecase, authRepo, formatResponse)

	e.Logger.Fatal(e.Start(":1323"))
}
