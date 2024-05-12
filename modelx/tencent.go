package modelx

import (
	"cloud.google.com/go/civil"
)

const (
	SourceKTBMainVault = "KTB_MAIN_VAULT"
	SourceGOVMainVault = "GOV_MAIN_VAULT"
	SourceKTBCIDVault  = "KTB_CID_VAULT"
	SourceGOVCIDVault  = "GOV_CID_VAULT"
	SourceTempVault    = "TEMP_VAULT"
	SourcePaotangVault = "PAOTANG_VAULT"
	SourceCompareImage = "COMPARE_IMAGE"

	CompareTypeReflect = "REFLECT"
	CompareTypeBasic   = "BASIC"
)

var (
	EngineResponse = map[int]string{
		0:     "Success",
		-1101: "Face detections failed",
		-1102: "Image decode failed",
		-1103: "Feature extraction failed",
		-1012: "Color sequence generation failed",
		-1301: "Invalid parameter",
		-1415: "Reflection data timestamp verification failed",
		-1490: "Time mismatch of reflection data",
		-4003: "Model initialization failed",
		-4005: "Similarity calculations failed",
		-4007: "Face identity cross check fails",
		-4014: "The system cannot detect face keypoints",
		-5005: "Face image extractions from the video failed",
		-5009: "Face Detection from the video failed.",
		-5011: "Liveness Detection from the video failed",
		-5017: "Liveness Detection from image reshooting analysis failed",
		-5018: "Liveness Detection from action recognition failed",
		-5019: "Reflections liveness detections failed",
		-5023: "Environmental light is too string",
		-5024: "Reflection detection core functions call failed",
	}
)

type FaceCompareHeader struct {
	RefID           string `json:"refId"`
	ReqID           string `json:"reqId"`
	CallbackURL     string `json:"callbackUrl"`
	Channel         string `json:"channel"`
	Product         string `json:"product"`
	TrustedSource   string `json:"trustedSource"`
	CompareType     string `json:"compareType"`
	SkipCompareFace bool   `json:"skipCompareFace"`
	PairId          string `json:"pairId"`
	SessionId       string `json:"sessionId"`
	StreamKafka     string `json:"streamKafka"`
	AdditionalInfo  string `json:"additionalInfo"`
}

type FaceCompareRequest struct {
	Header *FaceCompareHeader `json:"header"`
	Data   string             `json:"data"`
}

type FaceCompareResponse struct {
	OCRCallbackResponse *OCRCallbackResponse `json:"ocr,omitempty"`
	Identifier          string               `json:"identifier"`
	IdentifierType      string               `json:"identifierType"`
	TencentCode         int                  `json:"tencentCode"`
	TencentDesc         string               `json:"tencentDesc"`
	CompareCode         int                  `json:"compareCode"`
	CompareMsg          string               `json:"compareMsg"`
	PictureLiveCode     int                  `json:"pictureLiveCode"`
	PictureLiveMsg      string               `json:"pictureLiveMsg"`
	ReflectLiveCode     int                  `json:"reflectLiveCode"`
	ReflectLiveMsg      string               `json:"reflectLiveMsg"`
	Confidence          float64              `json:"confidence"`
	Result              bool                 `json:"result"`
	FaceCompareDate     string               `json:"faceCompareDate"`
}

type FaceComparisonHistory struct {
	RefID                        string         `json:"refId" bigquery:"ref_id"`
	ReqID                        string         `json:"reqId" bigquery:"req_id"`
	IdentifierEncrypted          string         `json:"identifierEncrypted" bigquery:"identifier_encrypted"`
	IdentifierType               string         `json:"identifierType" bigquery:"identifier_type"`
	ImageSelfie                  string         `json:"imageSelfie" bigquery:"image_selfie"`
	ImageCompare                 string         `json:"imageCompare" bigquery:"image_compare"`
	Channel                      string         `json:"channel" bigquery:"channel"`
	Product                      string         `json:"product" bigquery:"product"`
	ConfidenceLevel              float64        `json:"confidenceLevel" bigquery:"confidence_level"`
	Result                       bool           `json:"result" bigquery:"result"`
	TrustedSource                string         `json:"trustedSource" bigquery:"trusted_source"`
	CompareType                  string         `json:"compareType" bigquery:"compare_type"`
	CompareCode                  int            `json:"compareCode" bigquery:"compare_code"`
	CompareMsg                   string         `json:"compareMsg" bigquery:"compare_msg"`
	PictureLiveCode              int            `json:"pictureLiveCode" bigquery:"picture_live_code"`
	PictureLiveMsg               string         `json:"pictureLiveMsg" bigquery:"picture_live_msg"`
	ReflectLiveCode              int            `json:"reflectLiveCode" bigquery:"reflect_live_code"`
	ReflectLiveMsg               string         `json:"reflectLiveMsg" bigquery:"reflect_live_msg"`
	ErrorCode                    int            `json:"errorCode" bigquery:"error_code"`
	ErrorMSG                     string         `json:"errorMSG" bigquery:"error_msg"`
	OcrCode                      string         `json:"ocrCode" bigquery:"ocr_code"`
	OcrDescription               string         `json:"ocrDescription" bigquery:"ocr_description"`
	OcrDataInnoStatus            int            `json:"ocrDataInnoStatus" bigquery:"ocr_data_inno_status"`
	OcrDataInnoStatusDescription string         `json:"ocrDataInnoStatusDescription" bigquery:"ocr_data_inno_status_description"`
	Remark                       string         `json:"remark" bigquery:"remark"`
	Vault                        string         `json:"vault" bigquery:"-"`
	SaveVault                    bool           `json:"saveVault" bigquery:"-"`
	CallbackURL                  string         `json:"callbackUrl" bigquery:"callback_url"`
	RequestData                  string         `json:"requestData" bigquery:"request_data"`
	CreatedAt                    civil.DateTime `json:"createdAt" bigquery:"created_at"`
	Ref1                         string         `json:"ref1" bigquery:"ref1"`
	Ref2                         string         `json:"ref2" bigquery:"ref2"`
	Ref3                         string         `json:"ref3" bigquery:"ref3"`
	Ref4                         string         `json:"ref4" bigquery:"ref4"`
	Ref5                         string         `json:"ref5" bigquery:"ref5"`
	PairId                       string         `json:"pairId" bigquery:"pair_id"`
	SessionId                    string         `json:"sessionId" bigquery:"session_id"`
	StreamKafka                  string         `json:"streamKafka" bigquery:"-"`
}
