package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/config"
)

type contextKey string
type JwtCustomClaims struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
	jwt.StandardClaims
}

const UserKey contextKey = "userID"

func SignJwtAuthenticationToken(sub int) (string, error) {
	fmt.Println(config.ConfigInstance.JWT_TOKEN_SECRET)
	expiration := time.Now().Add(time.Duration(config.ConfigInstance.JWTExpirationInSeconds) * time.Second)
	claims := JwtCustomClaims{
		Sub: strconv.Itoa(sub),
		Iat: time.Now().Unix(),
		Exp: expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ConfigInstance.JWT_TOKEN_SECRET))
	return tokenString, err
}

func ValidateJwtAuthenticationToken(tokenString string) (*JwtCustomClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("authentication token is required")
	}
	claims := &JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigInstance.JWT_TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("access token expired")
	}

	return claims, err
}
