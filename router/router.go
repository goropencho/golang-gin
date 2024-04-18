package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goropencho/golang-gin/controller"
)

func NewRouter(userController *controller.UserController) *gin.Engine {

	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome User!")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	router := service.Group("/api")
	userRouter := router.Group("/users")
	userRouter.POST("", userController.Create)
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:id", userController.FindById)
	userRouter.PUT("/:id", userController.Update)
	userRouter.DELETE("/:id", userController.Delete)

	return service
}
