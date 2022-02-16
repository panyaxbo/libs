package errorx

import (
	"fmt"
	"testing"

	"github.com/panyaxbo/libs/utilx"
	"github.com/stretchr/testify/assert"
)

func TestMockTimeoutError(t *testing.T) {
	err := mockTimeout()
	if err != nil {
		fmt.Printf("%s", err.Error())
		assert.True(t, true)
	}

}
func TestCheckMockTimeoutError(t *testing.T) {
	err := mockTimeout()
	if utilx.IsErrorTimeout(err) {
		fmt.Printf("YES it's error timeout")
		assert.True(t, true)
	}

}
func mockTimeout() error {
	return TimeoutError{}
}
