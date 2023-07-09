package str

import (
	"strings"
)

func ExtractNumberFromText(source string) string {
	var result strings.Builder
	for i := 0; i < len(source); i++ {
		b := source[i]
		if '0' <= b && b <= '9' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func ExtractAlfaNumericFromText(source string) string {
	var result strings.Builder
	for i := 0; i < len(source); i++ {
		b := source[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

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
