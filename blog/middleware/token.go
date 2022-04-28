package middleware

import (
	"blog/config"
	pkg "blog/pkg"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ID string
	jwt.StandardClaims
}

func Parse(c *gin.Context) {
	//dvar stu controller.Json
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		pkg.Send(c, 401, "请求头中的auth为空")
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		pkg.Send(c, 401, "请求头中的auth格式有误")
		c.Abort()
		return
	}

	token, err := ParseToken(authHeader)
	if err != nil {
		pkg.Send(c, 401, "token无效")
		c.Abort()
		return
	}

	issuer := config.Issuer
	//_, err := model.GetUserInfoFormOne()

	if issuer != config.Issuer {
		pkg.Send(c, 401, "发布者无效")
		c.Abort()
		return
	}

	id := token.ID
	c.Set("id", id)
}
