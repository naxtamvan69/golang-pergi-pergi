package model

import "time"

type Destinasi struct {
	ID              int             `gorm:"primaryKey" json:"id"`
	NegaraDestinasi string          `gorm:"type:varchar(50);not null" json:"negara_destinasi"`
	IsBebasVisa     int             `gorm:"type:integer; not null" json:"is_bebas_visa"`
	AgensiTravels   []*TravelAgensi `gorm:"many2many:destinasi_travelagensi"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdateAt        time.Time       `json:"update_at"`
}
