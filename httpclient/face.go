package httpclient

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"
	"go-checkin/models"
	"os"
	"strconv"
	"time"
)

type HTTPClient struct {
	client *resty.Client
}

func NewHTTPClient() *HTTPClient {
	client := resty.New()
	if os.Getenv("APP_ENV") != "production" {
		client.SetDebug(true)
	}

	defaultTimeout, _ := strconv.Atoi(os.Getenv("DEFAULT_TIMEOUT"))
	client.SetTimeout(time.Second * time.Duration(defaultTimeout))

	return &HTTPClient{client: client}
}

func (hc *HTTPClient) CheckPhotoScore(image1 string, image2 string) (models.FaceCompare, error) {
	res, err := hc.client.R().
		SetFormData(map[string]string{
			"api_key":    "9vAJhzMUqMwvy-t2IuPsnMbqW_989vEJ",
			"api_secret": "dnLxTqE54o3-cg1KeYZRaVkvmLkM7Ho7",
		}).
		SetFiles(map[string]string{
			"image_file1": image1,
			"image_file2": image2,
		}).
		Post("https://api-us.faceplusplus.com/facepp/v3/compare")
	if err != nil {
		log.Error(err)
	}

	var response models.FaceCompare
	_ = json.Unmarshal(res.Body(), &response)

	return response, err
}
