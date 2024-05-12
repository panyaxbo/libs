package modelx

type CheckByLaserRequest struct {
	Header      RequestHeader `json:"header"`
	CitizenID   string        `json:"citizenId"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	DateOfBirth string        `json:"dateOfBirth"`
	LaserCode   string        `json:"laserCode"`
}

func (c *CheckByLaserRequest) Validate() (*Response, error) {
	if len(c.Header.Product) == 0 {
		return NewReponseBadRequest("Header Error: X-Product is required")
	}

	if len(c.Header.RefID) == 0 {
		return NewReponseBadRequest("Header Error: X-Ref-Id is required")
	}

	if len(c.CitizenID) == 0 {
		return NewReponseBadRequest("Body Error: citizenId is required")
	}

	if len(c.FirstName) == 0 {
		return NewReponseBadRequest("Body Error: firstName is required")
	}

	if len(c.LastName) == 0 {
		return NewReponseBadRequest("Body Error: lastName is required")
	}

	if len(c.DateOfBirth) == 0 {
		return NewReponseBadRequest("Body Error: dateOfBirth is required")
	}

	if len(c.LaserCode) == 0 {
		return NewReponseBadRequest("Body Error: laserCode is required")
	}

	return nil, nil
}

type DopaResult struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
