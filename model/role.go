package model

import "time"

type Role struct {
	ID        int       `json:"id" gorm:"primaryKey" binding:"required"`
	Role      string    `json:"role" gorm:"type:varchar(50);not null;unique" binding:"required"`
	UserRole  []User    `json:"user_role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdateAt  time.Time `json:"update_at" binding:"required"`
}

type RoleRequest struct {
	ID   int    `json:"id" binding:"required"`
	Role string `json:"role" binding:"required"`
}
