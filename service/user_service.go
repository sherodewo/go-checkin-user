package service

import (
	"go-checkin/dto"
	"go-checkin/models"
	"go-checkin/repository"
	"go-checkin/utils"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repository,
	}
}

func (s *UserService) FindAllUsers() (*[]models.User, error) {
	data, err := s.UserRepository.FindAll()
	return &data, err
}

func (s *UserService) FindUserById(id string) (*models.User, error) {
	data, err := s.UserRepository.FindById(id)

	return &data, err
}

func (s *UserService) SaveUser(dto dto.UserDto) (*models.User, error) {
	hashPassword, _ := utils.HashPassword(dto.Password)

	entity := models.User{
		Nik:        dto.Nik,
		Name:       dto.Name,
		Email:      dto.Email,
		Password:   hashPassword,
		IsActive:   dto.IsActive,
		TypeUser:   dto.TypeUser,
		UserRoleID: dto.UserRoleID,
		CreatedAt:  time.Now(),
	}
	data, err := s.UserRepository.Save(entity)
	return &data, err
}

func (s *UserService) UpdateUser(id string, dto dto.UserUpdateDto) (*models.User, error) {

	img := models.Photo{
		UserID:    id,
		Path:      dto.Image,
		CreatedAt: time.Now(),
	}
	dataImg, _ := s.UserRepository.UploadPhoto(img)

	password, _ := utils.HashPassword(dto.Password)

	var entity models.User
	entity.UserID = id
	entity.Name = dto.Name
	entity.Email = dto.Email
	entity.PhotoID = dataImg.ID

	if dto.Password != "" {
		entity.Password = password
	}
	//if fileName != "" {
	//	entity.ImageUrl = fileName
	//}
	data, err := s.UserRepository.Update(entity)

	return &data, err
}

func (s *UserService) DeleteUser(id string) error {
	entity := models.User{
		UserID: id,
	}
	err := s.UserRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *UserService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.User, err error) {
	recordTotal, err = s.UserRepository.Count()

	if searchValue != "" {
		recordFiltered, err = s.UserRepository.CountWhere("or", map[string]interface{}{
			"email LIKE ?": "%" + searchValue + "%",
			"nik LIKE ?":   "%" + searchValue + "%",
		})

		data, err = s.UserRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
			"email LIKE ?": "%" + searchValue + "%",
			"nik LIKE ?":   "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}

	recordFiltered, err = s.UserRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.UserRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}

func (s *UserService) GetDbInstance() *gorm.DB {
	return s.UserRepository.DbInstance()
}
