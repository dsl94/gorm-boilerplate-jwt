package user

import "github.com/jinzhu/gorm"

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []User {
	var users []User
	u.DB.Find(&users)

	return users
}

func (u *UserRepository) FindById(id uint) User {
	var user User
	u.DB.Find(&user, id)

	return user
}

func (u *UserRepository) Save(user User) User {
	u.DB.Save(&user)

	return user
}

func (u *UserRepository) Delete(user User) {
	u.DB.Delete(&user)
}

func (u *UserRepository) FindByUsername(username string) User {
	var user User
	u.DB.Where(&User{Username: username}).Preload("Roles").First(&user)

	return user
}
