package utilx

import (
	"encoding/json"
	"strings"
	"unicode/utf8"

	"github.com/panyaxbo/libs/configx"
	"github.com/panyaxbo/libs/maskx"
)

type SensitiveHeader struct {
	XIdentifierType string `json:"identifierType"`
	XIdentifier     string `json:"identifier"`
	XCitizenId      string `json:"citizenId"`
	XPassport       string `json:"passport"`
}

var (
	sensitiveHeaderFields = []string{
		"citizenId",
		"passport",
	}
	sensitiveXIdentifierTypes = []string{
		"citizen-id",
		"passport",
	}
)

func MaskHeaderWithEncryption(h SensitiveHeader, config *configx.ConfigMaskLog) SensitiveHeader {
	b, _ := json.Marshal(h)
	shf := getSensitiveHeaderFields(h.XIdentifierType)

	m := maskx.Init(shf)
	t, _ := m.JsonMaskEncrypted(b, config.Env)

	mh := SensitiveHeader{}
	json.Unmarshal([]byte(*t), &mh)

	return mh
}

func MaskHeaderWithSymbol(h SensitiveHeader, config *configx.ConfigMaskLog) SensitiveHeader {
	mh := SensitiveHeader{
		XIdentifierType: h.XIdentifierType,
		XIdentifier:     h.XIdentifier,
		XCitizenId:      strings.Repeat(config.Symbol, utf8.RuneCountInString(h.XCitizenId)),
		XPassport:       strings.Repeat(config.Symbol, utf8.RuneCountInString(h.XPassport)),
	}

	if doMaskXIdentifier(h.XIdentifierType) {
		mh.XIdentifier = strings.Repeat(config.Symbol, utf8.RuneCountInString(h.XIdentifier))
	}

	return mh
}

func getSensitiveHeaderFields(xIdentifierType string) []string {
	shf := sensitiveHeaderFields

	if doMaskXIdentifier(xIdentifierType) {
		shf = append(shf, "identifier")
	}

	return shf
}

func doMaskXIdentifier(xIdentifierType string) bool {
	return xIdentifierType == "" || isSensitiveXIdentifierType(xIdentifierType)
}

func isSensitiveXIdentifierType(x string) bool {
	return ContainsInSliceString(sensitiveXIdentifierTypes, x)
}
