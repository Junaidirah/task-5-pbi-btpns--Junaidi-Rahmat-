package router

import (
	"golang-api/controller"
	"golang-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.PUT("/:userId", controller.UserUpdate)
		userRouter.DELETE("/:userId", controller.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.ListPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controller.DeletePhoto)
	}

	return r
}
