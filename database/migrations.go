package database

import (
	"find-table/role"
	"find-table/user"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&role.Role{}, &user.User{})
}
