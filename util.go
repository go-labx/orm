package orm

import "strings"

func mapToString(m map[string]string) string {
	var str string
	for key, value := range m {
		str += key + "=" + value + "&"
	}
	return strings.TrimRight(str, "&")
}
