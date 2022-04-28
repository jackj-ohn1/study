package middleware

import (
	"blog/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(id string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	claims := Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Issuer,
			NotBefore: time.Now().Unix(), // 生效时间
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Key))
	return token, err
}

func ParseToken(tokenstring string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Key), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
