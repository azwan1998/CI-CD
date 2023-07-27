package appMiddleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	IdUser int
	Role   string
}

func CreateToken(personId int, role string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["personId"] = personId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))

}

func ExtractTokenUserId(c echo.Context) UserInfo {
	token := c.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		personId := claims["personId"]
		role := claims["role"]
		switch personId.(type) {
		case float64:
			return UserInfo{IdUser: int(personId.(float64)), Role: fmt.Sprintf("%v", role)}
		default:
			return UserInfo{IdUser: personId.(int), Role: fmt.Sprintf("%v", role)}
		}
	}
	return UserInfo{IdUser: -1, Role: ""} // invalid user
}
