package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/enescedev/gotodo/config"
	"github.com/enescedev/gotodo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

/*
func Jwt(c *gin.Context) {

		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
		   c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		   return
		}
		//compare the user from the request, with the one we defined:
		if user.Username != u.Username || user.Password != u.Password {
		   c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		   return
		}
		token, err := CreateToken(user.ID)
		if err != nil {
		   c.JSON(http.StatusUnprocessableEntity, err.Error())
		   return
		}
		c.JSON(http.StatusOK, token)
	  }
*/
type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetJWT(c *gin.Context) {

	var user User
	user.Id = 1
	user.Name = "Bulut Bilisimciler"
	// github.com/golang-jwt/jwt
	// TODO: Token üret...
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour).Unix(),
		"sub":  user.Id,
	})
	key := config.Key
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)

	fmt.Println(tokenString, err)
	// TODO: Token'ı kullanıcıya gönder...

	// middlewares/middleware_jwt.go (Yazılacak)
	c.JSON(200, gin.H{

		"token": (tokenString),
	})
}

func GetToDo(c *gin.Context) {
	todos := []models.Todo{}
	config.DB.Find(&todos)
	c.JSON(200, &todos)
}

func CreateToDo(c *gin.Context) {
	user, _ := c.Get("user")
	usermodel, _ := user.(User)
	log.Println("user=", usermodel)
	var todos models.Todo
	c.BindJSON(&todos)
	config.DB.Create(&todos)
	c.JSON(200, &todos)
}

func DeleteToDo(c *gin.Context) {
	var todos models.Todo
	config.DB.Where("id = ?", c.Param("id")).Delete(&todos)
	c.JSON(200, &todos)

}

func UpdateToDo(c *gin.Context) {
	var todos models.Todo
	config.DB.Where("id = ?", c.Param("id")).First(&todos)
	c.BindJSON(&todos)
	config.DB.Save(&todos)
	c.JSON(200, &todos)

}
