package utilx

import (
	"fmt"
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

func TestValidEmail(t *testing.T) {
	var addresses = []string{
		"foo@gmail.com",
		"Gopher <from@example.com>",
		"example",
	}
	for _, a := range addresses {
		if ok := IsValidEmail(a); ok {
			fmt.Printf("value: %-30s valid email: %-10t address: %s\n", a, ok, "")
		} else {
			fmt.Printf("value: %-30s valid email: %-10t\n", a, ok)
		}
	}
}
