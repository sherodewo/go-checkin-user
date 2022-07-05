package models

type PhotoRequest struct {
	Name string `json:"name"`
}

type Checkin struct {
	Langitude string `json:"langitude"`
	Longitude string `json:"longitude"`
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
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
