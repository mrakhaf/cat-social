package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWT struct{}

type jwtCustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateToken(userId string) (tokenString string, err error) {

	claims := &jwtCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("secret"))

	return
}

func (j *JWT) GetAccessFromToken(token string) (userId string, err error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return
	}

	claims, ok := tokenClaims.Claims.(*jwtCustomClaims)
	if !ok || !tokenClaims.Valid {
		err = echo.ErrUnauthorized
		return
	}

	userId = claims.UserId

	return
}
