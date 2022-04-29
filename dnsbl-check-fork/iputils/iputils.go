package iputils

import "strings"

func reverseStringByToken(str string, delimiter string) string {
	if str == "" {
		return ""
	}

	// strings.Cut not available in Go 1.16 (as set by the go.mod).
	var before, after string
	if i := strings.Index(str, delimiter); i >= 0 {
		before = str[:i]
		after = str[i+len(delimiter):]
	} else {
		return str
	}

	return reverseStringByToken(after, delimiter) + delimiter + before
}

// ReverseIPv4Address will return ip address octets in reverse order:
// 1.2.3.4 will be 4.3.2.1
func ReverseIPv4Address(str string) string {
	return reverseStringByToken(str, ".")
}
