package user

func RegisterToUser(register UserRegister) User {
	return User{Username: register.Username, FullName: register.FullName, Password: register.Password, Email: register.Email}
}

func UserToDto(user User) UserDto {
	return UserDto{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email}
}
