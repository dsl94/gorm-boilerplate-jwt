package user

import "find-table/role"

type UserService struct {
	UserRepository UserRepository
	RoleRepository role.RoleRepository
}

func ProvideUserService(u UserRepository, r role.RoleRepository) UserService {
	return UserService{UserRepository: u, RoleRepository: r}
}

func (u *UserService) FindAll() []UserDto {
	users := u.UserRepository.FindAll()
	var dtos = make([]UserDto, len(users))

	for i, user := range users {
		dtos[i] = UserToDto(user)
	}

	return dtos
}

func (u *UserService) FindById(id uint) User {
	return u.UserRepository.FindById(id)
}

func (u *UserService) Register(user User) User {
	adminRole := u.RoleRepository.FindByRole("ROLE_ADMIN")

	user.Roles = []role.Role{adminRole}

	u.UserRepository.Save(user)

	return user
}

func (u *UserService) Delete(user User) {
	u.UserRepository.Delete(user)
}

func (u *UserService) FindByUsername(username string) User {
	return u.UserRepository.FindByUsername(username)
}

func (u *UserService) CheckLoginCredentials(login UserLogin) (User, bool) {
	user := u.UserRepository.FindByUsername(login.Username)

	err := user.CheckPassword(login.Password)
	if err != nil {
		return User{}, false
	}

	return user, true
}
