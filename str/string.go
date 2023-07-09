package str

import (
	"strings"
)

func SplitOnNotEmpty(str string, delimiter string) []string {
	if strings.TrimSpace(str) != "" {
		var result []string
		for _, item := range strings.Split(str, delimiter) {
			if item != "" && item != " " {
				result = append(result, item)
			}
		}
		return result
	}
	return nil
}

func RemoveUncommonCharacters(str string) string {
	newBody := ""
	for _, c := range str {
		if c > 160 {
			continue
		}
		newBody += string(c)
	}
	return strings.TrimSpace(newBody)
}
