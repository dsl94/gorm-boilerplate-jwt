package role

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RoleController struct {
	RoleService RoleService
}

func ProvideRoleController(r RoleService) RoleController {
	return RoleController{RoleService: r}
}

func (r *RoleController) Create(c *gin.Context) {

	var roleDto RoleDto
	err := c.BindJSON(&roleDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdRole := r.RoleService.Save(ToRole(roleDto))

	c.JSON(http.StatusOK, gin.H{"role": ToDto(createdRole)})
}
