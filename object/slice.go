package object

func SliceContains(data []string, target string) bool {
	for _, s := range data {
		if target == s {
			return true
		}
	}
	return false
}

func SliceNotContains(data []string, target string) bool {
	return !SliceContains(data, target)
}
