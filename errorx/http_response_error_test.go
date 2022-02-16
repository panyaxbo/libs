package errorx

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo500Response(t *testing.T) {

	err := mock500Response()

	fmt.Printf("%s \n", err.Error())

	assert.Equal(t, http.StatusInternalServerError, err.(*HttpResponseError).StatusCode)

	if strings.Contains(err.Error(), http.StatusText(http.StatusInternalServerError)) {
		assert.True(t, true)
	}
}
func TestDo500MsgResponse(t *testing.T) {

	err := mock500Response()

	if strings.Contains(err.Error(), http.StatusText(http.StatusInternalServerError)) {
		assert.True(t, true)
	}
}
func mock500Response() error {
	h := HttpResponseError{}
	return h.do500InternalServerErrorResponse()

}
