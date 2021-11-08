package utilx

import (
	"encoding/json"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/panyaxbo/libs/logx"
)

const (
	escapeChar = "\b\f\v\r\t\n"
)

func RemoveEscapeCharactor(str string) string {
	return strings.TrimSuffix(strings.TrimRight(strings.TrimSpace(str), escapeChar), escapeChar)
}

func StructToString(data interface{}) string {
	d, err := json.Marshal(data)
	if err != nil {
		logx.Errorf("Cannot Convert Struct -> String %s", err)
		return ""
	}
	return string(d)
}

func IsStringNullBlank(str string) bool {
	return str == "" || len(str) <= 0
}

func IsErrorTimeout(err error) bool {
	istimeout := false
	switch err := err.(type) {
	case net.Error:
		if err.Timeout() {
			istimeout = true
		}
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			istimeout = true
		}
	}
	if strings.Contains(err.Error(), "Client.Timeout") {
		istimeout = true
	}
	if os.IsTimeout(err) {
		istimeout = true
	}
	return istimeout
}
