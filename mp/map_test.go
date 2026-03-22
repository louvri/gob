package mp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Copy ---

func TestCopy_IgnoresSpecifiedKeys(t *testing.T) {
	m := map[string]any{"a": 1, "b": 2, "c": 3}
	result := Copy(m, []string{"b"})
	assert.Equal(t, map[string]any{"a": 1, "c": 3}, result)
}

func TestCopy_EmptyIgnore(t *testing.T) {
	m := map[string]any{"a": 1, "b": 2}
	result := Copy(m, []string{})
	assert.Equal(t, map[string]any{"a": 1, "b": 2}, result)
}

func TestCopy_IgnoreAll(t *testing.T) {
	m := map[string]any{"a": 1, "b": 2}
	result := Copy(m, []string{"a", "b"})
	assert.Empty(t, result)
}

func TestCopy_EmptyMap(t *testing.T) {
	m := map[string]any{}
	result := Copy(m, []string{"a"})
	assert.Empty(t, result)
}

func TestCopy_IntKeys(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	result := Copy(m, []int{2})
	assert.Equal(t, map[int]string{1: "a", 3: "c"}, result)
}

// --- CopyOnly ---

func TestCopyOnly_FiltersSpecifiedKeys(t *testing.T) {
	m := map[string]any{"a": 1, "b": 2, "c": 3}
	result := CopyOnly(m, []string{"a", "c"})
	assert.Equal(t, map[string]any{"a": 1, "c": 3}, result)
}

func TestCopyOnly_EmptyFilter(t *testing.T) {
	m := map[string]any{"a": 1, "b": 2}
	result := CopyOnly(m, []string{})
	assert.Empty(t, result)
}

func TestCopyOnly_FilterNotPresent(t *testing.T) {
	m := map[string]any{"a": 1}
	result := CopyOnly(m, []string{"z"})
	assert.Empty(t, result)
}

func TestCopyOnly_EmptyMap(t *testing.T) {
	m := map[string]any{}
	result := CopyOnly(m, []string{"a"})
	assert.Empty(t, result)
}

// --- Search ---

func TestSearch_FindsFirstExistingKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	key, found := Search(m, []string{"z", "b", "c"})
	assert.True(t, found)
	assert.Equal(t, "b", key)
}

func TestSearch_NoneFound(t *testing.T) {
	m := map[string]int{"a": 1}
	key, found := Search(m, []string{"x", "y"})
	assert.False(t, found)
	assert.Equal(t, "", key)
}

func TestSearch_EmptyKeys(t *testing.T) {
	m := map[string]int{"a": 1}
	key, found := Search(m, []string{})
	assert.False(t, found)
	assert.Equal(t, "", key)
}

func TestSearch_EmptyMap(t *testing.T) {
	m := map[string]int{}
	key, found := Search(m, []string{"a"})
	assert.False(t, found)
	assert.Equal(t, "", key)
}

func TestSearch_IntKeys(t *testing.T) {
	m := map[int]string{10: "a", 20: "b"}
	key, found := Search(m, []int{5, 20})
	assert.True(t, found)
	assert.Equal(t, 20, key)
}
