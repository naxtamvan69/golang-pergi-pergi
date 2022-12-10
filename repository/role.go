package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"pergipergi/model"
)

type RoleRepository interface {
	AddRole(ctx context.Context, role model.Role) (model.Role, error)
	UpdateRole(ctx context.Context, role model.Role) (model.Role, error)
	DeleteRole(ctx context.Context, ID int) error
	GetRoles(ctx context.Context) ([]model.Role, error)
	GetRoleByID(ctx context.Context, ID int) (model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) AddRole(ctx context.Context, role model.Role) (model.Role, error) {
	err := r.db.WithContext(ctx).Model(&model.Role{}).Create(&role).Error
	if err != nil {
		return model.Role{}, err
	}
	return role, err
}

func (r *roleRepository) UpdateRole(ctx context.Context, role model.Role) (model.Role, error) {
	err := r.db.WithContext(ctx).Model(&model.Role{}).Where("id = ?", role.ID).Updates(&role).Error
	if err != nil {
		return model.Role{}, err
	}
	return role, nil
}

func (r *roleRepository) DeleteRole(ctx context.Context, ID int) error {
	err := r.db.WithContext(ctx).Where("id = ?", ID).Delete(model.Role{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) GetRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.WithContext(ctx).Model(&model.Role{}).Select("*").Scan(&roles).Error
	if err != nil {
		return nil, err
	} else if len(roles) == 0 {
		return nil, errors.New("roles data is empty")
	}
	return roles, nil
}

func (r *roleRepository) GetRoleByID(ctx context.Context, ID int) (model.Role, error) {
	var role model.Role
	err := r.db.WithContext(ctx).Model(&model.Role{}).Select("*").Where("id = ?", ID).Scan(&role).Error
	if err != nil {
		return model.Role{}, err
	} else if role.ID == 0 {
		return model.Role{}, errors.New("find user by id = " + string(ID) + " not found")
	}
	return role, err
}
