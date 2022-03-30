package utilx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidThaiNationalID(t *testing.T) {

	//assert.Equal(t, http.StatusBadRequest, err.(*HttpResponse400Error).StatusCode)

	a := IsValidThaiNationalID("3460100645535")
	assert.True(t, true, a)
}
func TestNotValidThaiNationalID(t *testing.T) {

	//assert.Equal(t, http.StatusBadRequest, err.(*HttpResponse400Error).StatusCode)

	a := IsValidThaiNationalID("111111111111")
	assert.False(t, false, a)

	b := IsValidThaiNationalID("AA343434343434")
	//	fmt.Printf("TestNotValidThaiNationalID b %t", b)
	assert.False(t, false, b)
}
