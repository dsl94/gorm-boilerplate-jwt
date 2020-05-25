package role

import "github.com/jinzhu/gorm"

type RoleRepository struct {
	DB *gorm.DB
}

func ProvideRoleRepository(DB *gorm.DB) RoleRepository {
	return RoleRepository{DB : DB}
}

func (r *RoleRepository) Save(role Role) Role {
	r.DB.Save(&role)

	return role
}

func (r *RoleRepository) FindByRole(roleName string) Role {
	var role Role
	r.DB.Where(&Role{RoleName: roleName}).First(&role)

	return role
}