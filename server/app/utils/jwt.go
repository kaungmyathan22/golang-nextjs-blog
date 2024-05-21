package jwt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/config"
)

type contextKey string

const UserKey contextKey = "userID"

func SignJwtAuthenticationToken(sub int) (string, error) {
	fmt.Println(config.ConfigInstance.JWT_TOKEN_SECRET)
	expiration := time.Second * time.Duration(config.ConfigInstance.JWTExpirationInSeconds)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		ID:        strconv.Itoa(sub),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "golang-nextjs-blog",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ConfigInstance.JWT_TOKEN_SECRET))
	return tokenString, err
}

func ValidateJwtAuthenticationToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.ConfigInstance.JWT_TOKEN_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return &claims, nil
	}
	return nil, fmt.Errorf("invalid token")

}
