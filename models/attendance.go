package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PhotoRequest struct {
	Name string `json:"name"`
}

type Checkin struct {
	ID        string `json:"id" form:"id"`
	Maps      string `json:"maps" form:"maps" validate:"required"`
	ProjectID string `json:"project_id" form:"project_id" validate:"required"`
	ImageUrl  string `json:"image_url" form:"image_url" validate:"required"`
}

type FaceCompare struct {
	RequestID  string  `json:"request_id"`
	TimeUsed   int     `json:"time_used"`
	Confidence float64 `json:"confidence"`
	Thresholds struct {
		OneE3 float64 `json:"1e-3"`
		OneE4 float64 `json:"1e-4"`
		OneE5 float64 `json:"1e-5"`
	} `json:"thresholds"`
	Faces1 []struct {
		FaceToken     string `json:"face_token"`
		FaceRectangle struct {
			Top    int `json:"top"`
			Left   int `json:"left"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"face_rectangle"`
	} `json:"faces1"`
	Faces2 []struct {
		FaceToken     string `json:"face_token"`
		FaceRectangle struct {
			Top    int `json:"top"`
			Left   int `json:"left"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"face_rectangle"`
	} `json:"faces2"`
	ImageID1 string `json:"image_id1"`
	ImageID2 string `json:"image_id2"`
}

type Maps struct {
	Info struct {
		Statuscode int `json:"statuscode"`
		Copyright  struct {
			Text         string `json:"text"`
			ImageURL     string `json:"imageUrl"`
			ImageAltText string `json:"imageAltText"`
		} `json:"copyright"`
		Messages []interface{} `json:"messages"`
	} `json:"info"`
	Options struct {
		MaxResults        int  `json:"maxResults"`
		ThumbMaps         bool `json:"thumbMaps"`
		IgnoreLatLngInput bool `json:"ignoreLatLngInput"`
	} `json:"options"`
	Results []struct {
		ProvidedLocation struct {
			LatLng struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"latLng"`
		} `json:"providedLocation"`
		Locations []struct {
			Street             string `json:"street"`
			AdminArea6         string `json:"adminArea6"`
			AdminArea6Type     string `json:"adminArea6Type"`
			AdminArea5         string `json:"adminArea5"`
			AdminArea5Type     string `json:"adminArea5Type"`
			AdminArea4         string `json:"adminArea4"`
			AdminArea4Type     string `json:"adminArea4Type"`
			AdminArea3         string `json:"adminArea3"`
			AdminArea3Type     string `json:"adminArea3Type"`
			AdminArea1         string `json:"adminArea1"`
			AdminArea1Type     string `json:"adminArea1Type"`
			PostalCode         string `json:"postalCode"`
			GeocodeQualityCode string `json:"geocodeQualityCode"`
			GeocodeQuality     string `json:"geocodeQuality"`
			DragPoint          bool   `json:"dragPoint"`
			SideOfStreet       string `json:"sideOfStreet"`
			LinkID             string `json:"linkId"`
			UnknownInput       string `json:"unknownInput"`
			Type               string `json:"type"`
			LatLng             struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"latLng"`
			DisplayLatLng struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"displayLatLng"`
		} `json:"locations"`
	} `json:"results"`
}

type Presence struct {
	ID           string    `gorm:"column:id;primary_key:true"`
	UserID       string    `gorm:"column:user_id"`
	Checkin      time.Time `gorm:"column:checkin"`
	Checkout     time.Time `gorm:"column:checkout"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	PhotoID      string    `gorm:"column:photo_id"`
	LocationName string    `gorm:"column:location_name"`
	Location     string    `gorm:"column:location"`
}

func (c *Presence) TableName() string {
	return "presence"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Presence) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}

type AttendDatatable struct {
	ID        string `json:"id"`
	ImagePath string `json:"image_path"`
	Day       string `json:"day"`
	Hours     string `json:"hours"`
	IN        string `json:"in"`
	Out       string `json:"out"`
}
