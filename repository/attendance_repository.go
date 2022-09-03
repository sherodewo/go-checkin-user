package repository

import (
	"go-checkin/models"
	"gorm.io/gorm"
	"time"
)

type AttendanceRepository interface {
	Save(presence models.Presence) error
	SavePhoto(photo models.Photo) (models.Photo, error)
	Update(models.Checkin) error
	Checkout(id string) error
	GetImageByUserID(userID string) (models.Photo, error)
	GetPhotoID(photoID string) (models.Photo, error)
	QueryDatatable(id string) ([]models.Presence, error, int64)
}

type attendanceRepository struct {
	*gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{DB: db}
}

func (r *attendanceRepository) Save(checkin models.Presence) error {
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

func (r *attendanceRepository) Checkout(id string) error {
	err := r.DB.Model(models.Presence{ID: id}).UpdateColumns(&models.Presence{Checkout: time.Now()}).Error
	return err
}

func (r *attendanceRepository) GetImageByUserID(userID string) (models.Photo, error) {
	var entity models.Photo
	err := r.DB.Where("user_id = ?", userID).Model(models.Photo{}).Find(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *attendanceRepository) GetPhotoID(photoID string) (models.Photo, error) {
	var entity models.Photo
	err := r.DB.Where("id = ?", photoID).Model(models.Photo{}).Find(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *attendanceRepository) SavePhoto(entity models.Photo) (models.Photo, error) {
	err := r.DB.Create(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *attendanceRepository) QueryDatatable(id string) ([]models.Presence, error, int64) {
	var entity []models.Presence
	err := r.DB.
		Model(models.Presence{}).
		Where("user_id = ?", id).
		Limit(10).
		Order("created_at desc").
		Find(&entity).Error
	if err != nil {
		return entity, err, 0
	}
	newErr := r.DB.
		Model(models.Presence{}).
		Where("user_id = ?", id).
		Order("created_at desc").
		Find(&entity)
	return entity, nil, newErr.RowsAffected
}
