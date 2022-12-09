package repository

import (
	"context"
	"gorm.io/gorm"
	"pergipergi/model"
)

type DestinasiRepository interface {
	AddDestinasi(ctx context.Context, destinasi model.Destinasi) (model.Destinasi, error)
	UpdateDestinasi(ctx context.Context, destinasi model.Destinasi) (model.Destinasi, error)
	DeleteDestinasi(ctx context.Context, ID int) error
	GetDestinations(ctx context.Context) ([]model.Destinasi, error)
	GetDestinasiByNegaraDestinasi(ctx context.Context, negara string) (model.Destinasi, error)
}

type destinasiRepository struct {
	db *gorm.DB
}

func NewDestinasiRepository(db *gorm.DB) *destinasiRepository {
	return &destinasiRepository{db}
}

func (d *destinasiRepository) AddDestinasi(ctx context.Context, destinasi model.Destinasi) (model.Destinasi, error) {
	err := d.db.WithContext(ctx).Model(&model.Destinasi{}).Create(&destinasi).Error
	if err != nil {
		return model.Destinasi{}, err
	}
	return destinasi, nil
}

func (d *destinasiRepository) UpdateDestinasi(ctx context.Context, destinasi model.Destinasi) (model.Destinasi, error) {
	err := d.db.WithContext(ctx).Model(&model.Destinasi{}).Where("id = ?", destinasi.ID).Updates(&destinasi).Error
	if err != nil {
		return model.Destinasi{}, err
	}
	return destinasi, nil
}

func (d *destinasiRepository) DeleteDestinasi(ctx context.Context, ID int) error {
	err := d.db.WithContext(ctx).Where("id = ?", ID).Delete(model.Destinasi{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *destinasiRepository) GetDestinations(ctx context.Context) ([]model.Destinasi, error) {
	var destinations []model.Destinasi
	err := d.db.WithContext(ctx).Model(&model.Destinasi{}).Select("*").Scan(&destinations).Error
	if err != nil {
		return nil, err
	}
	return destinations, nil
}

func (d *destinasiRepository) GetDestinasiByNegaraDestinasi(ctx context.Context, negara string) (model.Destinasi, error) {
	var destinasi model.Destinasi
	err := d.db.WithContext(ctx).Model(&model.Destinasi{}).Select("*").Where("negara_destinasi = ?", negara).Scan(&destinasi).Error
	if err != nil {
		return model.Destinasi{}, err
	}
	return destinasi, nil
}
