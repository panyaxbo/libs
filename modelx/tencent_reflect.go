package modelx

import "encoding/json"

type ReflectCompareRequest struct {
	Identifier         string             `json:"identifier"`
	IdentifierType     string             `json:"identifierType"`
	DiscardVault       bool               `json:"discardVault"`
	CheckOCR           bool               `json:"checkOCR"`
	Reference          *Reference         `json:"reference,omitempty"`
	ReflectCompareData ReflectCompareData `json:"reflectCompareData"`
}

func (r *ReflectCompareRequest) Validate(header *FaceCompareHeader) (*Response, error) {
	if header.Product == "" {
		return NewReponseBadRequest("X-Product is required")
	}

	if header.RefID == "" {
		return NewReponseBadRequest("X-Ref-Id is required")
	}

	if r.Identifier == "" {
		return NewReponseBadRequest("identifier is required")
	}

	if r.IdentifierType == "" {
		return NewReponseBadRequest("identifierType is required")
	}

	return nil, nil
}

type ReflectCompareData struct {
	ColorSequence string      `json:"color_data"`
	ReflectData   ReflectData `json:"reflect_data"`
	LiveImage     string      `json:"live_image"`
	Platform      int         `json:"platform"`
	CompareImage  string      `json:"compare_image"`
}

type ReflectData struct {
	OffsetSys       json.Number `json:"offset_sys"`
	Height          int         `json:"height"`
	LandmarkNum     int         `json:"landmark_num"`
	ImagesData      ImagesData  `json:"images_data"`
	Width           int         `json:"width"`
	ConfigBegin     int         `json:"config_begin"`
	Log             string      `json:"log"`
	BeginTime       int64       `json:"begin_time"`
	ChangepointTime int64       `json:"changepoint_time"`
	FrameNum        int         `json:"frame_num"`
}

type ImagesData []struct {
	CaptureTime int64  `json:"capture_time"`
	Image       string `json:"image"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
}

func (r *ReflectCompareRequest) ReplaceCompareImage(newCompareImage string) {
	r.ReflectCompareData.CompareImage = newCompareImage
}
