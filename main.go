package main

import (
	"fmt"
	"strings"
)

func maliciousIP(adr string) (bool, error) {
	if !isIPv4(adr) {
		return false, fmt.Errorf("only ipv4 addresses supported")
	}

	return true, nil
}

func isIPv4(adr string) bool {
	return strings.Count(adr, ":") < 2
}
