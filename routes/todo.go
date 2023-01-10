package routes

import (
	"github.com/enescedev/gotodo/controller"
	"github.com/enescedev/gotodo/middleware"
	"github.com/gin-gonic/gin"
)

func ToDoRoute(router *gin.Engine) {
	router.Use(middleware.JwtRequiredMiddleware)
	router.GET("/generate-jwt", controller.GetJWT)
	router.GET("/", controller.GetToDo)
	router.POST("/createtodo", middleware.JwtRequiredMiddleware, controller.CreateToDo)
	router.DELETE("/:id", controller.DeleteToDo)
	router.PUT("/:id", controller.UpdateToDo)
}
