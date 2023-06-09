package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string, isAdmin bool) string {
	// menyimpan data user didalam token
	claims := jwt.MapClaims{
		"id":      id,
		"email":   email,
		"isAdmin": isAdmin,
	}

	// menentukan metode enkripsi yg digunakan
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// membuat token yang sudah ditanda tangani oleh secret key
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("gagal membuat token")
	}

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
