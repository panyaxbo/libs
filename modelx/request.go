package modelx

import "net/http"

type RequestHeader struct {
	RefID          string `json:"refID"`
	CallbackURL    string `json:"callbackURL"`
	Channel        string `json:"channel"`
	Product        string `json:"product"`
	Source         string `json:"source"`
	APIKey         string `json:"apiKey"`
	StreamKafka    string `json:"streamKafka"`
	AdditionalInfo string `json:"additionalInfo"`
}

func NewRequestHeader(headers http.Header) RequestHeader {
	return RequestHeader{
		RefID:          headers.Get(HeaderXRefID),
		CallbackURL:    headers.Get(HeaderXCallbackURL),
		Channel:        headers.Get(HeaderXChannel),
		Product:        headers.Get(HeaderXProduct),
		Source:         headers.Get(HeaderXSource),
		APIKey:         headers.Get(HeaderXApiKey),
		StreamKafka:    headers.Get(HeaderXStreamKafka),
		AdditionalInfo: headers.Get(HeaderXAdditionalInfo),
	}
}
