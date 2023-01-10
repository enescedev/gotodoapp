package middleware

import (
	"fmt"
	"net/http"

	"github.com/enescedev/gotodo/config"
	"github.com/enescedev/gotodo/controller"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtRequiredMiddleware(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Header Yok!"})
		return
	}

	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		key := config.Key
		return key, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["exp"], claims["user"])
		var user controller.User
		dto, _ := claims["user"].(map[string]interface{})
		user.Id = int64(dto["id"].(float64))
		user.Name = dto["name"].(string)
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "JWT Invalid Değil!"})
		fmt.Println(err)
		return
	}

}
