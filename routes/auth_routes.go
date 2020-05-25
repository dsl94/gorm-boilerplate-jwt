package routes

import (
	"find-table/role"
	"find-table/user"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, roleController *role.RoleController, userController *user.UserController) {
	router.POST("/api/roles", roleController.Create)
	router.POST("/api/register", userController.Register)
}
