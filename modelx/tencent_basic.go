package modelx

type BasicCompareRequest struct {
	Identifier       string           `json:"identifier"`
	IdentifierType   string           `json:"identifierType"`
	DiscardVault     bool             `json:"discardVault"`
	Reference        *Reference       `json:"reference,omitempty"`
	BasicCompareData BasicCompareData `json:"basicCompareData"`
}

type BasicCompareData struct {
	SelfieImage  string `json:"selfieImage"`
	CompareImage string `json:"compareImage"`
}

func (b *BasicCompareRequest) Validate(header *FaceCompareHeader) (*Response, error) {
	if header.Product == "" {
		return NewReponseBadRequest("X-Product is required")
	}

	if header.RefID == "" {
		return NewReponseBadRequest("X-Ref-Id is required")
	}

	if b.Identifier == "" {
		return NewReponseBadRequest("identifier is required")
	}

	if b.IdentifierType == "" {
		return NewReponseBadRequest("identifierType is required")
	}

	if b.BasicCompareData.SelfieImage == "" {
		return NewReponseBadRequest("selfieImage is required")
	}

	return nil, nil
}

func (b *BasicCompareRequest) ReplaceCompareImage(newCompareImage string) {
	b.BasicCompareData.CompareImage = newCompareImage
}
