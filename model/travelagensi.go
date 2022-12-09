package model

import "time"

type TravelAgensi struct {
	ID           int          `gorm:"primaryKey" json:"id"`
	NamaAgensi   string       `gorm:"type:varchar(30);not null" json:"nama_agensi"`
	AlamatAgensi string       `gorm:"type:varchar(255);not null" json:"alamat_agensi"`
	NomorTelepon string       `gorm:"type:varchar(13);not null" json:"nomor_telepon"`
	WaktuBuka    time.Time    `gorm:"not null" json:"waktu_buka"`
	WaktuTutup   time.Time    `gorm:"not null" json:"waktu_tutup"`
	Destinations []*Destinasi `gorm:"many2many:destinasi_travelagensi"`
	TourGuides   []User       `gorm:"foreignKey:TravelAgensiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdateAt     time.Time    `json:"update_at"`
}
