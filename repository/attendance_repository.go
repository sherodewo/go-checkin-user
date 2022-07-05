package repository

import (
	"go-checkin/models"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Save(models.Checkin) error
	Update(models.Checkin) error
	GetImageByUserID(userID string) (models.Photo, error)
}

type attendanceRepository struct {
	*gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{DB: db}
}

func (r *attendanceRepository) Save(checkin models.Checkin) error {
	err := r.DB.Create(&checkin).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *attendanceRepository) Update(checkin models.Checkin) error {
	err := r.DB.Create(&checkin).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *attendanceRepository) GetImageByUserID(userID string) (models.Photo, error) {
	var entity models.Photo
	err := r.DB.Model(models.Photo{ID: userID}).Find(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}
