package utilx

import "strings"

func RemoveEscapeCharactor(str string) string {
	return strings.TrimRight(strings.TrimSpace(str), "\r\t\n")
}
