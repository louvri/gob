package mp

func Copy[K comparable, V any](m map[K]V, ignore []K) map[K]V {
	ignoreSet := make(map[K]bool, len(ignore))
	for _, k := range ignore {
		ignoreSet[k] = true
	}
	result := make(map[K]V)
	for key, value := range m {
		if !ignoreSet[key] {
			result[key] = value
		}
	}
	return result
}

func CopyOnly[K comparable, V any](m map[K]V, filter []K) map[K]V {
	filterSet := make(map[K]bool, len(filter))
	for _, k := range filter {
		filterSet[k] = true
	}
	result := make(map[K]V)
	for key, value := range m {
		if filterSet[key] {
			result[key] = value
		}
	}
	return result
}

func Search[K comparable, V any](data map[K]V, keys []K) (K, bool) {
	for _, k := range keys {
		if _, ok := data[k]; ok {
			return k, true
		}
	}
	var zero K
	return zero, false
}
