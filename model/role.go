package model

import "time"

type Role struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Role      string    `json:"role" gorm:"type:varchar(50);not null;unique"`
	UserRole  []User    `json:"user_role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
