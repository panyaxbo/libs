package utilx

import "strings"

const (
	escapeChar = "\b\f\v\r\t\n"
)

func RemoveEscapeCharactor(str string) string {
	return strings.TrimSuffix(strings.TrimRight(strings.TrimSpace(str), escapeChar), escapeChar)
}
