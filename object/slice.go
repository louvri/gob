package object

// Deprecated: SliceContains is superseded by arr.Search. Use arr.Search(data, target) != -1 instead.
func SliceContains(data []string, target string) bool {
	for _, s := range data {
		if target == s {
			return true
		}
	}
	return false
}

// Deprecated: SliceNotContains is superseded by arr.Search. Use arr.Search(data, target) == -1 instead.
func SliceNotContains(data []string, target string) bool {
	return !SliceContains(data, target)
}
