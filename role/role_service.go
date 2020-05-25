package role

type RoleService struct {
	RoleRepository RoleRepository
}

func ProvideRoleService(r RoleRepository) RoleService {
	return RoleService{RoleRepository: r}
}

func (r *RoleService) Save(role Role) Role {
	r.RoleRepository.Save(role)

	return role
}

func (r *RoleService) FindByRoleName(name string) Role {
	role := r.RoleRepository.FindByRole(name)

	return role
}