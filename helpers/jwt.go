package helpers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"

	"github.com/gin-gonic/gin"
)

var secretKey = "your-256-bit-secret"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		// "exp":   time.Now().Add(time.Hour * 3),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errReponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errReponse
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errReponse
		}
		return []byte(secretKey), nil
	})
	fmt.Println(token)
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errReponse
	}
	fmt.Println(token.Claims.(jwt.MapClaims))
	return token.Claims.(jwt.MapClaims), nil
}
