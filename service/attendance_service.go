package service

import (
	"go-checkin/httpclient"
	"go-checkin/models"
	"go-checkin/repository"
)

type AttendanceService struct {
	AttendanceRepository repository.AttendanceRepository
}

func NewAttendanceService(repository repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{
		AttendanceRepository: repository,
	}
}

func (s *AttendanceService) Save(req models.Checkin) error {
	err := s.AttendanceRepository.Save(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *AttendanceService) Update(req models.Checkin) error {
	err := s.AttendanceRepository.Update(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *AttendanceService) PhotoCompare(req models.PhotoRequest, userID string) (bool, error) {

	// Get Image by user ID
	image, err := s.AttendanceRepository.GetImageByUserID(userID)
	if err != nil {
		return false, err
	}
	httpClient := httpclient.NewHTTPClient()
	resp, err := httpClient.CheckPhotoScore(req.Name, image.Path)
	if err != nil {
		return false, err
	}
	if resp.Confidence < 50 {
		return false, err
	}
	return true, nil
}
