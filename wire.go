//+build wireinject

package main

import (
	"find-table/role"
	"find-table/user"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitRoleController(db *gorm.DB) role.RoleController {
	wire.Build(role.ProvideRoleRepository, role.ProvideRoleService, role.ProvideRoleController)

	return role.RoleController{}
}

func InitUserController(db *gorm.DB, repository role.RoleRepository) user.UserController {
	wire.Build(user.ProvideUserRepository, user.ProvideUserService, user.ProvideUserController)

	return user.UserController{}
}
