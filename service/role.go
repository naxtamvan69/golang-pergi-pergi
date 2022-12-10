package service

import (
	"context"
	"pergipergi/model"
	"pergipergi/repository"
)

type RoleService interface {
	AddRoleService(ctx context.Context, role model.Role) (model.Role, error)
	UpdateRoleService(ctx context.Context, role model.Role) (model.Role, error)
	DeleteRoleService(ctx context.Context, ID int) error
	GetRolesService(ctx context.Context) ([]model.Role, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *roleService {
	return &roleService{roleRepository}
}

func (r *roleService) AddRoleService(ctx context.Context, role model.Role) (model.Role, error) {
	return r.roleRepository.AddRole(ctx, role)
}

func (r *roleService) UpdateRoleService(ctx context.Context, role model.Role) (model.Role, error) {
	_, err := r.roleRepository.GetRoleByID(ctx, role.ID)
	if err != nil {
		return model.Role{}, err
	}
	return r.roleRepository.UpdateRole(ctx, role)
}

func (r *roleService) DeleteRoleService(ctx context.Context, ID int) error {
	_, err := r.roleRepository.GetRoleByID(ctx, ID)
	if err != nil {
		return err
	}
	return r.roleRepository.DeleteRole(ctx, ID)
}

func (r *roleService) GetRolesService(ctx context.Context) ([]model.Role, error) {
	return r.roleRepository.GetRoles(ctx)
}
