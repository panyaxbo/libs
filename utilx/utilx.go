package utilx

import (
	"encoding/json"
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
