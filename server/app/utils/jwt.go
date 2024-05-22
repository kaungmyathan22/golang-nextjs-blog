package jwt

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
	expiration := time.Now().Add(1 * time.Hour)
	claims := JwtCustomClaims{
		Sub: strconv.Itoa(sub),
		Iat: time.Now().Unix(),
		Exp: expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	claims := JwtCustomClaims{
		Exp:       jwt.NewNumericDate(time.Now().Add(expiration)),
		ID:        strconv.Itoa(sub),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "golang-nextjs-blog",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ConfigInstance.JWT_TOKEN_SECRET))
	return tokenString, err
}

func ValidateJwtAuthenticationToken(tokenString string) (*jwt.RegisteredClaims, error) {
	fmt.Println(tokenString)
	fmt.Println("tokenString")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigInstance.JWT_TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("error while parsing token to register claims")
	}
	return &claims, err
}
