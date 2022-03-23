package errorx

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttp500Response(t *testing.T) {

	err := mockHttp500Response()

	fmt.Printf("%s \n", err.Error())

	assert.Equal(t, http.StatusInternalServerError, err.(*HttpResponse500Error).StatusCode)

	if strings.Contains(err.Error(), http.StatusText(http.StatusInternalServerError)) {
		assert.True(t, true)
	}
}

func mockHttp500Response() error {
	h := HttpResponse500Error{}
	return h.InternalServerError500Response()
}
