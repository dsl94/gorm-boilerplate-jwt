package user

import (
	"errors"
	"find-table/role"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string
	FullName string
	Email    string
	Password string
	Roles    []role.Role `gorm:"many2many:user_role;"`
}

func (user *User) HashPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be null")
	}
	bytePassword := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	hashedPassword := []byte(user.Password)

	return bcrypt.CompareHashAndPassword(hashedPassword, bytePassword)
}
