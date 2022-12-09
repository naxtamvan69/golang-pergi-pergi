package repository

import (
	"context"
	"gorm.io/gorm"
	"pergipergi/model"
)

type TravelAgensiRepository interface {
	AddTravelAgensi(ctx context.Context, travelAgensi model.TravelAgensi) (model.TravelAgensi, error)
	UpdateTravelAgensi(ctx context.Context, travelAgensi model.TravelAgensi) (model.TravelAgensi, error)
	DeleteTravelAgensi(ctx context.Context, ID int) error
	GetListTravelAgensi(ctx context.Context) ([]model.TravelAgensi, error)
	GetTravelAgensiByID(ctx context.Context, ID int) (model.TravelAgensi, error)
}

type travelAgensiRepository struct {
	db *gorm.DB
}

func NewTravelAgensiRepository(db *gorm.DB) *travelAgensiRepository {
	return &travelAgensiRepository{db}
}

func (t *travelAgensiRepository) AddTravelAgensi(ctx context.Context, travelAgensi model.TravelAgensi) (model.TravelAgensi, error) {
	err := t.db.WithContext(ctx).Model(&model.TravelAgensi{}).Create(&travelAgensi).Error
	if err != nil {
		return model.TravelAgensi{}, err
	}
	return travelAgensi, nil
}

func (t *travelAgensiRepository) UpdateTravelAgensi(ctx context.Context, travelAgensi model.TravelAgensi) (model.TravelAgensi, error) {
	err := t.db.WithContext(ctx).Model(&model.TravelAgensi{}).Where("id = ?", travelAgensi.ID).Updates(&travelAgensi).Error
	if err != nil {
		return model.TravelAgensi{}, err
	}
	return travelAgensi, nil
}

func (t *travelAgensiRepository) DeleteTravelAgensi(ctx context.Context, ID int) error {
	err := t.db.WithContext(ctx).Where("id = ?", ID).Delete(model.Destinasi{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *travelAgensiRepository) GetListTravelAgensi(ctx context.Context) ([]model.TravelAgensi, error) {
	var listTravelAgensi []model.TravelAgensi
	err := t.db.WithContext(ctx).Model(&model.TravelAgensi{}).Select("*").Scan(&listTravelAgensi).Error
	if err != nil {
		return nil, err
	}
	return listTravelAgensi, nil
}

func (t *travelAgensiRepository) GetTravelAgensiByID(ctx context.Context, ID int) (model.TravelAgensi, error) {
	var travelAgensi model.TravelAgensi
	err := t.db.WithContext(ctx).Select("*").Where("id = ?", ID).Scan(&travelAgensi).Error
	if err != nil {
		return model.TravelAgensi{}, err
	}
	return travelAgensi, nil
}
