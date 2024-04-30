package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authHandler "github.com/mrakhaf/cat-social/domain/auth/delivery/http"
	authRepo "github.com/mrakhaf/cat-social/domain/auth/repository"
	authUsecase "github.com/mrakhaf/cat-social/domain/auth/usecase"
	catHandler "github.com/mrakhaf/cat-social/domain/cat/delivery/http"
	catRepo "github.com/mrakhaf/cat-social/domain/cat/repository"
	catUsecase "github.com/mrakhaf/cat-social/domain/cat/usecase"
	"github.com/mrakhaf/cat-social/shared/common"
	formatJson "github.com/mrakhaf/cat-social/shared/common/json"
	"github.com/mrakhaf/cat-social/shared/common/jwt"
	"github.com/mrakhaf/cat-social/shared/config/database"
	"github.com/mrakhaf/cat-social/shared/config/setup"
)

func main() {
	e := echo.New()

	setup.SetupTimezone()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = common.NewValidator()

	err := godotenv.Load("conf/config.env")
	if err != nil {
		e.Logger.Fatal(err)
	}

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

	catGroup := e.Group("/v1/cat")
	catGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	}))

	formatResponse := formatJson.NewResponse()
	jwtAccess := jwt.NewJWT()

	//auth
	authRepo := authRepo.NewRepository(catDB)
	authUsecase := authUsecase.NewUsecase(authRepo, jwtAccess)
	authHandler.AuthHandler(userGroup, authUsecase, authRepo, formatResponse)

	//cat
	catRepo := catRepo.NewRepository(catDB)
	catUsecase := catUsecase.NewUsecase(catRepo)
	catHandler.CatHandler(catGroup, formatResponse, jwtAccess, catUsecase)

	e.Logger.Fatal(e.Start(":1323"))
}
