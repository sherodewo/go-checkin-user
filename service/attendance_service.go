package service

import (
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go-checkin/httpclient"
	"go-checkin/models"
	"go-checkin/repository"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

type AttendanceService struct {
	AttendanceRepository repository.AttendanceRepository
}

func NewAttendanceService(repository repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{
		AttendanceRepository: repository,
	}
}

func (s *AttendanceService) Save(req models.Checkin, photo models.Photo) error {
	// Marshal maps info
	var maps models.Maps
	_ = json.Unmarshal([]byte(req.Maps), &maps)
	locName := maps.Results[0].Locations[0].Street + ", " + maps.Results[0].Locations[0].AdminArea5 + ", " + maps.Results[0].Locations[0].AdminArea3
	payload, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(maps)

	attendance := models.Presence{
		UserID:       req.ID,
		Checkin:      time.Now(),
		CreatedAt:    time.Now(),
		PhotoID:      photo.ID,
		LocationName: locName,
		Location:     string(payload),
	}

	err := s.AttendanceRepository.Save(attendance)
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
func (s *AttendanceService) Checkout(id string) error {
	err := s.AttendanceRepository.Checkout(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AttendanceService) PhotoCompare(req models.Checkin) (bool, error, *models.Photo) {
	url := req.ImageUrl
	// Get Image by user ID
	image, err := s.AttendanceRepository.GetImageByUserID(req.ID)
	if err != nil {
		return false, err, nil
	}

	filename := path.Base(url)
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
		return false, err, nil
	}
	defer resp.Body.Close()

	name := "assets/avatar/" + filename + ".jpg"
	f, err := os.Create(name)
	if err != nil {
		return false, err, nil
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	fmt.Println("FILE 1 : ", name)
	fmt.Println("FILE 2 : ", image.Path)
	httpClient := httpclient.NewHTTPClient()
	compareResp, err := httpClient.CheckPhotoScore(name, image.Path)
	if err != nil {
		return false, err, nil
	}
	if compareResp.Confidence < 50 {
		return false, err, nil
	}
	//Set to DB photo
	photo := models.Photo{
		UserID:    req.ID,
		Path:      name,
		CreatedAt: time.Now(),
	}
	svPhoto, err := s.AttendanceRepository.SavePhoto(photo)
	if err != nil {
		return false, errors.New("failed save to DB"), nil
	}
	return true, nil, &svPhoto
}

func (s *AttendanceService) QueryDatatable(id string) ([]models.AttendDatatable, error, int64) {
	var resp []models.AttendDatatable
	var val string
	var Out string
	data, err, total := s.AttendanceRepository.QueryDatatable(id)
	if err != nil {
		return nil, err, total
	}

	for _, v := range data {
		//Get Image Path
		photo, _ := s.AttendanceRepository.GetPhotoID(v.PhotoID)

		timeIn := v.Checkin
		timeOut := v.Checkout

		Day := timeIn.Format("Monday ,02 January")
		In := timeIn.Format(" 15:04 ")
		total := timeOut.Sub(timeIn)

		if timeOut.After(timeIn) {
			val = fmt.Sprintf("Worked for %d Hrs %d Min", int64(total.Hours()), int64(total.Minutes())-(int64(total.Hours())*60))
			Out = timeOut.Format(" 15:04 ")
		} else {
			val = ""
			Out = "WORKING"
		}
		check := models.AttendDatatable{
			ID:        v.ID,
			ImagePath: photo.Path,
			Day:       Day,
			Hours:     val,
			IN:        In,
			Out:       Out,
		}
		resp = append(resp, check)
	}

	return resp, nil, total
}
