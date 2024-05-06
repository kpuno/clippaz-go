package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kpuno/clippaz/controllers"
	"github.com/kpuno/clippaz/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("user")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
