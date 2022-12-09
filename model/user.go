package model

import "time"

type User struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	Username       string `gorm:"type:varchar(50);not null;unique" json:"username"`
	Name           string `gorm:"type:varchar(50);not null" json:"name"`
	Age            int    `gorm:"type:integer;not null" json:"age"`
	Sex            int    `gorm:"type:integer;not null" json:"sex"`
	Password       string `gorm:"type:varchar(255);not null" json:"password"`
	RoleID         int    `json:"role_id"`
	TravelAgensiID int
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"update_at"`
}
