package main

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	authHandler "github.com/mrakhaf/cat-social/domain/auth/delivery/http"
	"github.com/mrakhaf/cat-social/domain/auth/repository"
	"github.com/mrakhaf/cat-social/shared/config/database"
)

func main() {
	e := echo.New()

	//db config
	catDB, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(echojwt.JWT([]byte("secret")))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Cat Social!!!")
	})

	//create group
	group := e.Group("/v1")

	//auth
	repository.NewRepository(catDB)
	authHandler.AuthHandler(group, nil, nil)

	e.Logger.Fatal(e.Start(":1323"))
}
