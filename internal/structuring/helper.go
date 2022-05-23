package structuring

import (
	"strings"
)

// toUpFirstletter In string up first letter
func toUpFirstletter(str string) string {
	return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
}

// canAddToAppend work with array and items. If you need unique items in array - using this method
// p.s If you have type string and []string
func canAddToAppend(value string, arr []string) bool {
	ok := true
	for _, newImport := range arr {
		if newImport == value {
			ok = false
		}
	}

	return ok
}
