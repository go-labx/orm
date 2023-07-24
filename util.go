package orm

import (
	"bytes"
	"strings"
	"unicode"
)

// MapToString is a utility function that converts a map into a string.
// Each key-value pair in the map is converted into a "key=value" string, and these strings are joined with "&" as the separator.
// The resulting string does not end with "&".
func MapToString(m map[string]string) string {
	var str string
	for key, value := range m {
		str += key + "=" + value + "&"
	}
	return strings.TrimRight(str, "&")
}

// ToSnakeCase is a utility function that converts a given string into snake case.
func ToSnakeCase(s string) string {
	var buffer bytes.Buffer
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}
