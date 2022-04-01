package utilx

import (
	"encoding/json"
	"net"
	"net/mail"
	"net/url"
	"os"
	"regexp"
	"strconv"
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

func IsValidThaiNationalID(id string) bool {
	matched, _ := regexp.MatchString(`^[0-9]{13}`, id)
	if !matched {
		return false
	}
	var sum = 0
	for i := 0; i < 12; i++ {
		v, _ := strconv.Atoi(id[i : i+1])
		sum += v * (13 - i)
	}
	var checkSum = (11 - sum%11) % 10
	charAt12, _ := strconv.Atoi(id[12:13])
	if checkSum == charAt12 {
		return true
	}
	return false
}

func IsValidEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}
