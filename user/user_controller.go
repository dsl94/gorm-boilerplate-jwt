package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	UserService UserService
}

func ProvideUserController(u UserService) UserController {
	return UserController{UserService: u}
}

func (u *UserController) Register(c *gin.Context) {
	var userRegister UserRegister
	err := c.BindJSON(&userRegister)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	user := RegisterToUser(userRegister)
	user.HashPassword(userRegister.Password)

	u.UserService.Register(user)

	c.Status(http.StatusOK)
}

func (u *UserController) FindAll(c *gin.Context) {
	//if !authorizer.AuthorizeRequestForRoles([]string{"ROLE_ADMIN"}, c) {
	//	return
	//}
	users := u.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": users})
}
