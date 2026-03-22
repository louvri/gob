package arr

func Search[T comparable](array []T, search T) int {
	for i, item := range array {
		if item == search {
			return i
		}
	}
	return -1
}

func Insert[T any](array []T, data T, index int) []T {
	if index < 0 || index > len(array) {
		panic("arr.Insert: index out of range")
	}
	result := make([]T, 0, len(array)+1)
	result = append(result, array[:index]...)
	result = append(result, data)
	result = append(result, array[index:]...)
	return result
}

func Copy[T comparable](array []T, ignore []T) []T {
	ignoreSet := make(map[T]bool, len(ignore))
	for _, item := range ignore {
		ignoreSet[item] = true
	}
	result := make([]T, 0)
	for _, item := range array {
		if !ignoreSet[item] {
			result = append(result, item)
		}
	}
	return result
}

func Trim(array []string) []string {
	result := make([]string, 0)
	for _, item := range array {
		if item != "" && item != " " {
			result = append(result, item)
		}
	}
	return result
}

func Unique[T comparable](arr1 []T, arr2 []T) []T {
	result := make([]T, 0)
	index1 := make(map[T]bool)
	index2 := make(map[T]bool)
	for _, item := range arr1 {
		index1[item] = true
	}
	for _, item := range arr2 {
		if !index1[item] {
			result = append(result, item)
		}
		index2[item] = true
	}
	for _, item := range arr1 {
		if !index2[item] {
			result = append(result, item)
		}
	}
	return result
}

// Deprecated: UniqueInt is superseded by Unique[int64]. Use that instead.
func UniqueInt(arr1 []int64, arr2 []int64) []int64 {
	return Unique(arr1, arr2)
}

func Index[T comparable](columns []T) map[T]bool {
	idx := make(map[T]bool, len(columns))
	for _, c := range columns {
		idx[c] = true
	}
	return idx
}

func Map[T any, R any](data []T, fn func(T) R) []R {
	result := make([]R, len(data))
	for i, item := range data {
		result[i] = fn(item)
	}
	return result
}

func Filter[T any](data []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func Reduce[T any, R any](data []T, initial R, fn func(R, T) R) R {
	acc := initial
	for _, item := range data {
		acc = fn(acc, item)
	}
	return acc
}
