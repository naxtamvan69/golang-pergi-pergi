package model

import "time"

type User struct {
	ID             int       `gorm:"primaryKey" json:"id" binding:"required"`
	Username       string    `gorm:"type:varchar(50);not null;unique" json:"username" binding:"required"`
	Name           string    `gorm:"type:varchar(50);not null" json:"name" binding:"required"`
	Age            int       `gorm:"type:integer;not null" json:"age" binding:"required"`
	Sex            int       `gorm:"type:integer;not null" json:"sex" binding:"required"`
	Password       string    `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
	RoleID         *int      `json:"role_id"`
	TravelAgensiID *int      `json:"travel_agensi_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"update_at"`
}

type UserRequest struct {
	ID             int    `json:"id" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Age            int    `json:"age" binding:"required"`
	Sex            int    `json:"sex" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RoleID         *int   `json:"role_id"`
	TravelAgensiID *int   `json:"travel_agensi_id"`
}
