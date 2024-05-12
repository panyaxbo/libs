package modelx

type OCRRequest struct {
	Base64 string `json:"base64"`
}

type OCRResponse struct {
	Status    int    `json:"status"`
	StatusMsg string `json:"status_msg"`
	Ocr       struct {
		Result struct {
			Field []struct {
				Name      string `json:"name"`
				OcrResult struct {
					Value string `json:"value"`
				} `json:"ocr_result"`
			} `json:"field"`
		} `json:"result"`
	} `json:"ocr"`
}

type OCRCallbackResponse struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Data        OCRCallbackData `json:"data"`
}

type OCRCallbackData struct {
	Status     int    `json:"status"`
	StatusMsg  string `json:"statusMsg"`
	Identifier string `json:"identifier"`
}
