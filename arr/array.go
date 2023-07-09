package arr

func Search(array []string, search string) int {
	for i, item := range array {
		if item == search {
			return i
		}
	}
	return -1
}

func Insert(array []string, data string, index int) []string {
	result := array[:index]
	tmp := array[index:]
	result = append(result, data)
	result = append(result, tmp...)
	return result
}

func Copy(array []string, ignore []string) []string {
	result := make([]string, 0)
	for _, item := range array {
		skip := false
		for _, search := range ignore {
			if search == item {
				skip = true
				break
			}
		}
		if !skip {
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
func Unique(arr1 []string, arr2 []string) []string {
	result := make([]string, 0)
	index1 := make(map[string]bool)
	index2 := make(map[string]bool)
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
func UniqueInt(arr1 []int64, arr2 []int64) []int64 {
	result := make([]int64, 0)
	index1 := make(map[int64]bool)
	index2 := make(map[int64]bool)
	for _, item := range arr1 {
		index1[item] = true
	}
	for _, item := range arr2 {
		if !index1[item] {
			result = append(result, item)
		}
	}
	for _, item := range arr1 {
		if !index2[item] {
			result = append(result, item)
		}
	}
	return result
}

func Index(columns []string) map[string]bool {
	idx := make(map[string]bool)
	for _, c := range columns {
		idx[c] = true
	}
	return idx
}
