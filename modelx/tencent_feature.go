package modelx

type FaceFeatureCompareResponse struct {
	ErrorCode       int     `json:"error_code"`
	ErrorMsg        string  `json:"error_msg"`
	Similarity      int     `json:"similarity"`
	SimilarityFloat float64 `json:"similarity_float"`
}

type FaceFeatureCompareRequest struct {
	Image   string `json:"image"`
	Feature string `json:"feature"`
}

type CompareFaceFeatureRequest struct {
	Identifier             string                    `json:"identifier"`
	IdentifierType         string                    `json:"identifierType"`
	FaceFeatureCompareData FaceFeatureCompareRequest `json:"faceFeatureCompareData"`
}

type CompareFaceFeatureResponse struct {
	TencentCode int     `json:"tencentCode"`
	TencentDesc string  `json:"tencentDesc"`
	ErrorCode   int     `json:"errorCode"`
	ErrorMsg    string  `json:"errorMSG"`
	Confidence  float64 `json:"confidence"`
	Result      bool    `json:"result"`
}
