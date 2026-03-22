package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Search ---

func TestSearch_Found(t *testing.T) {
	assert.Equal(t, 2, Search([]string{"a", "b", "c"}, "c"))
}

func TestSearch_NotFound(t *testing.T) {
	assert.Equal(t, -1, Search([]string{"a", "b", "c"}, "z"))
}

func TestSearch_EmptySlice(t *testing.T) {
	assert.Equal(t, -1, Search([]string{}, "a"))
}

func TestSearch_FirstElement(t *testing.T) {
	assert.Equal(t, 0, Search([]string{"x", "y"}, "x"))
}

func TestSearch_Int(t *testing.T) {
	assert.Equal(t, 1, Search([]int{10, 20, 30}, 20))
}

func TestSearch_IntNotFound(t *testing.T) {
	assert.Equal(t, -1, Search([]int{10, 20, 30}, 99))
}

// --- Insert ---

func TestInsert_Middle(t *testing.T) {
	result := Insert([]string{"a", "c", "d"}, "b", 1)
	assert.Equal(t, []string{"a", "b", "c", "d"}, result)
}

func TestInsert_Beginning(t *testing.T) {
	result := Insert([]string{"b", "c"}, "a", 0)
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestInsert_End(t *testing.T) {
	result := Insert([]string{"a", "b"}, "c", 2)
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestInsert_EmptySlice(t *testing.T) {
	result := Insert([]string{}, "a", 0)
	assert.Equal(t, []string{"a"}, result)
}

func TestInsert_DoesNotMutateOriginal(t *testing.T) {
	original := []string{"a", "b", "c"}
	copy := make([]string, len(original))
	for i, v := range original {
		copy[i] = v
	}
	Insert(original, "x", 1)
	assert.Equal(t, copy, original)
}

func TestInsert_PanicsOnNegativeIndex(t *testing.T) {
	assert.Panics(t, func() {
		Insert([]string{"a"}, "b", -1)
	})
}

func TestInsert_PanicsOnOutOfRange(t *testing.T) {
	assert.Panics(t, func() {
		Insert([]string{"a"}, "b", 5)
	})
}

// --- Copy ---

func TestCopy_IgnoreSome(t *testing.T) {
	result := Copy([]string{"a", "b", "c", "d"}, []string{"b", "d"})
	assert.Equal(t, []string{"a", "c"}, result)
}

func TestCopy_IgnoreNone(t *testing.T) {
	result := Copy([]string{"a", "b"}, []string{})
	assert.Equal(t, []string{"a", "b"}, result)
}

func TestCopy_IgnoreAll(t *testing.T) {
	result := Copy([]string{"a", "b"}, []string{"a", "b"})
	assert.Empty(t, result)
}

func TestCopy_EmptyArray(t *testing.T) {
	result := Copy([]string{}, []string{"a"})
	assert.Empty(t, result)
}

func TestCopy_Int(t *testing.T) {
	result := Copy([]int{1, 2, 3, 4}, []int{2, 4})
	assert.Equal(t, []int{1, 3}, result)
}

// --- Trim ---

func TestTrim_RemovesEmptyAndSpaces(t *testing.T) {
	result := Trim([]string{"a", "", " ", "b", "", "c"})
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestTrim_NoRemoval(t *testing.T) {
	result := Trim([]string{"a", "b"})
	assert.Equal(t, []string{"a", "b"}, result)
}

func TestTrim_AllEmpty(t *testing.T) {
	result := Trim([]string{"", " ", "", " "})
	assert.Empty(t, result)
}

func TestTrim_EmptySlice(t *testing.T) {
	result := Trim([]string{})
	assert.Empty(t, result)
}

// --- Unique ---

func TestUnique_SymmetricDifference(t *testing.T) {
	result := Unique([]string{"a", "b", "c"}, []string{"b", "c", "d"})
	assert.ElementsMatch(t, []string{"a", "d"}, result)
}

func TestUnique_NoOverlap(t *testing.T) {
	result := Unique([]string{"a", "b"}, []string{"c", "d"})
	assert.ElementsMatch(t, []string{"a", "b", "c", "d"}, result)
}

func TestUnique_CompleteOverlap(t *testing.T) {
	result := Unique([]string{"a", "b"}, []string{"a", "b"})
	assert.Empty(t, result)
}

func TestUnique_EmptyFirst(t *testing.T) {
	result := Unique([]string{}, []string{"a", "b"})
	assert.ElementsMatch(t, []string{"a", "b"}, result)
}

func TestUnique_EmptySecond(t *testing.T) {
	result := Unique([]string{"a", "b"}, []string{})
	assert.ElementsMatch(t, []string{"a", "b"}, result)
}

func TestUnique_BothEmpty(t *testing.T) {
	result := Unique([]string{}, []string{})
	assert.Empty(t, result)
}

func TestUnique_Int(t *testing.T) {
	result := Unique([]int{1, 2, 3}, []int{2, 3, 4})
	assert.ElementsMatch(t, []int{1, 4}, result)
}

// --- UniqueInt (deprecated wrapper) ---

func TestUniqueInt_Wrapper(t *testing.T) {
	result := UniqueInt([]int64{1, 2, 3}, []int64{2, 3, 4})
	assert.ElementsMatch(t, []int64{1, 4}, result)
}

// --- Index ---

func TestIndex_CreatesLookupMap(t *testing.T) {
	idx := Index([]string{"a", "b", "c"})
	assert.True(t, idx["a"])
	assert.True(t, idx["b"])
	assert.True(t, idx["c"])
	assert.False(t, idx["d"])
}

func TestIndex_Empty(t *testing.T) {
	idx := Index([]string{})
	assert.Empty(t, idx)
}

func TestIndex_Int(t *testing.T) {
	idx := Index([]int{1, 2, 3})
	assert.True(t, idx[1])
	assert.False(t, idx[99])
}

// --- Map ---

func TestMap_StringToLen(t *testing.T) {
	result := Map([]string{"a", "bb", "ccc"}, func(s string) int { return len(s) })
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestMap_IntDouble(t *testing.T) {
	result := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestMap_Empty(t *testing.T) {
	result := Map([]int{}, func(n int) int { return n })
	assert.Empty(t, result)
}

// --- Filter ---

func TestFilter_EvenNumbers(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestFilter_NoneMatch(t *testing.T) {
	result := Filter([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 })
	assert.Empty(t, result)
}

func TestFilter_AllMatch(t *testing.T) {
	result := Filter([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestFilter_Empty(t *testing.T) {
	result := Filter([]int{}, func(n int) bool { return true })
	assert.Empty(t, result)
}

// --- Reduce ---

func TestReduce_Sum(t *testing.T) {
	result := Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
	assert.Equal(t, 10, result)
}

func TestReduce_Concat(t *testing.T) {
	result := Reduce([]string{"a", "b", "c"}, "", func(acc, s string) string { return acc + s })
	assert.Equal(t, "abc", result)
}

func TestReduce_Empty(t *testing.T) {
	result := Reduce([]int{}, 42, func(acc, n int) int { return acc + n })
	assert.Equal(t, 42, result)
}
