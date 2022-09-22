package utilx

import (
	"testing"

	"github.com/panyaxbo/libs/configx"
	"github.com/panyaxbo/libs/utilx"
	"github.com/stretchr/testify/assert"
)

var (
	config = &configx.ConfigMaskLog{}
	env    = "dev"
)

func TestPrintHeaderWithMaskSymbol_Case1(t *testing.T) {
	config.WithMaskingLogWithEncrypt(env)

	h := utilx.SensitiveHeader{
		XIdentifierType: "citizen-id",
		XIdentifier:     "1234567890123",
		XCitizenId:      "1234567890123",
		XPassport:       "1234567890123",
	}

	mh := utilx.MaskHeaderWithEncryption(h, config)

	assert.Equal(t, "citizen-id", mh.XIdentifierType)
	assert.NotEqual(t, "1234567890123", mh.XIdentifier)
	assert.NotEqual(t, "1234567890123", mh.XCitizenId)
	assert.NotEqual(t, "1234567890123", mh.XPassport)
}

func TestPrintHeaderWithMaskSymbol_Case2(t *testing.T) {
	config.WithMaskingLogWithEncrypt(env)

	h := utilx.SensitiveHeader{
		XIdentifierType: "customer-key",
		XIdentifier:     "890000001",
		XCitizenId:      "1234567890123",
		XPassport:       "1234567890123",
	}

	mh := utilx.MaskHeaderWithEncryption(h, config)

	assert.Equal(t, "customer-key", mh.XIdentifierType)
	assert.Equal(t, "890000001", mh.XIdentifier)
	assert.NotEqual(t, "1234567890123", mh.XCitizenId)
	assert.NotEqual(t, "1234567890123", mh.XPassport)
}

func TestMaskHeaderWithSymbol_Case1(t *testing.T) {
	config.WithMaskingLogWithSymbol("*")

	h := utilx.SensitiveHeader{
		XIdentifierType: "citizen-id",
		XIdentifier:     "1234567890123",
		XCitizenId:      "1234567890123",
		XPassport:       "1234567890123",
	}

	mh := utilx.MaskHeaderWithSymbol(h, config)

	assert.Equal(t, "citizen-id", mh.XIdentifierType)
	assert.Equal(t, "*************", mh.XIdentifier)
	assert.Equal(t, "*************", mh.XCitizenId)
	assert.Equal(t, "*************", mh.XPassport)
}

func TestMaskHeaderWithSymbol_Case2(t *testing.T) {
	config.WithMaskingLogWithSymbol("*")

	h := utilx.SensitiveHeader{
		XIdentifierType: "customer-key",
		XIdentifier:     "890000001",
		XCitizenId:      "1234567890123",
		XPassport:       "1234567890123",
	}

	mh := utilx.MaskHeaderWithSymbol(h, config)

	assert.Equal(t, "customer-key", mh.XIdentifierType)
	assert.Equal(t, "890000001", mh.XIdentifier)
	assert.Equal(t, "*************", mh.XCitizenId)
	assert.Equal(t, "*************", mh.XPassport)
}
