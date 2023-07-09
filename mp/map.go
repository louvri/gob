package mp

func Copy(array map[string]interface{}, ignore []string) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range array {
		skip := false
		for _, search := range ignore {
			if key == search {
				skip = true
				break
			}
		}
		if !skip {
			result[key] = value
		}
	}
	return result
}
func CopyOnly(array map[string]interface{}, filter []string) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range array {
		for _, search := range filter {
			if key == search {
				result[key] = value
			}
		}
	}
	return result
}
func Search(data map[string]interface{}, key []string) string {
	for _, k := range key {
		if data[k] != nil {
			return k
		}
	}
	return ""
}
