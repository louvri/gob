package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- ExtractNumberFromText ---

func TestExtractNumberFromText_Mixed(t *testing.T) {
	assert.Equal(t, "123", ExtractNumberFromText("abc123def"))
}

func TestExtractNumberFromText_OnlyNumbers(t *testing.T) {
	assert.Equal(t, "456", ExtractNumberFromText("456"))
}

func TestExtractNumberFromText_NoNumbers(t *testing.T) {
	assert.Equal(t, "", ExtractNumberFromText("abcdef"))
}

func TestExtractNumberFromText_Empty(t *testing.T) {
	assert.Equal(t, "", ExtractNumberFromText(""))
}

func TestExtractNumberFromText_SpecialChars(t *testing.T) {
	assert.Equal(t, "12", ExtractNumberFromText("!@#1$%^2&*"))
}

// --- ExtractAlfaNumericFromText ---

func TestExtractAlfaNumericFromText_Mixed(t *testing.T) {
	assert.Equal(t, "Hello World 123", ExtractAlfaNumericFromText("Hello, World! 123"))
}

func TestExtractAlfaNumericFromText_OnlyAlphaNumeric(t *testing.T) {
	assert.Equal(t, "abc123", ExtractAlfaNumericFromText("abc123"))
}

func TestExtractAlfaNumericFromText_OnlySpecial(t *testing.T) {
	assert.Equal(t, "", ExtractAlfaNumericFromText("!@#$%"))
}

func TestExtractAlfaNumericFromText_Empty(t *testing.T) {
	assert.Equal(t, "", ExtractAlfaNumericFromText(""))
}

func TestExtractAlfaNumericFromText_PreservesSpaces(t *testing.T) {
	assert.Equal(t, "a b c", ExtractAlfaNumericFromText("a b c"))
}

// --- ExtractAlfaNumericWithSelectedSpecialCharactersFromText ---

func TestExtractAlfaNumericWithSelectedSpecialCharactersFromText_Original(t *testing.T) {
	sourceText := "Ab9 \\!@#$%^&*_+-=\n"
	validateText := "Ab9 \\!#$%&*+-\n"
	filteredText := ExtractAlfaNumericWithSelectedSpecialCharactersFromText(sourceText)
	assert.Equal(t, validateText, filteredText)
}

func TestExtractAlfaNumericWithSelectedSpecialCharactersFromText_Empty(t *testing.T) {
	assert.Equal(t, "", ExtractAlfaNumericWithSelectedSpecialCharactersFromText(""))
}

func TestExtractAlfaNumericWithSelectedSpecialCharactersFromText_AlphaOnly(t *testing.T) {
	assert.Equal(t, "Hello", ExtractAlfaNumericWithSelectedSpecialCharactersFromText("Hello"))
}

func TestExtractAlfaNumericWithSelectedSpecialCharactersFromText_NewlinePreserved(t *testing.T) {
	assert.Equal(t, "line1\nline2", ExtractAlfaNumericWithSelectedSpecialCharactersFromText("line1\nline2"))
}

// --- SplitOnNotEmpty ---

func TestSplitOnNotEmpty_Basic(t *testing.T) {
	result := SplitOnNotEmpty("a,b,c", ",")
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestSplitOnNotEmpty_WithEmptyParts(t *testing.T) {
	result := SplitOnNotEmpty("a,,b, ,c", ",")
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestSplitOnNotEmpty_EmptyString(t *testing.T) {
	result := SplitOnNotEmpty("", ",")
	assert.Nil(t, result)
}

func TestSplitOnNotEmpty_WhitespaceOnly(t *testing.T) {
	result := SplitOnNotEmpty("   ", ",")
	assert.Nil(t, result)
}

func TestSplitOnNotEmpty_SingleItem(t *testing.T) {
	result := SplitOnNotEmpty("hello", ",")
	assert.Equal(t, []string{"hello"}, result)
}

func TestSplitOnNotEmpty_DifferentDelimiter(t *testing.T) {
	result := SplitOnNotEmpty("a|b|c", "|")
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

// --- RemoveUncommonCharacters ---

func TestRemoveUncommonCharacters_NoUncommon(t *testing.T) {
	assert.Equal(t, "hello world", RemoveUncommonCharacters("hello world"))
}

func TestRemoveUncommonCharacters_WithUnicode(t *testing.T) {
	// Characters > 160 should be removed
	result := RemoveUncommonCharacters("hello\u00A1world")
	assert.Equal(t, "helloworld", result)
}

func TestRemoveUncommonCharacters_Empty(t *testing.T) {
	assert.Equal(t, "", RemoveUncommonCharacters(""))
}

func TestRemoveUncommonCharacters_TrimSpaces(t *testing.T) {
	assert.Equal(t, "hello", RemoveUncommonCharacters("  hello  "))
}

func TestRemoveUncommonCharacters_AllUncommon(t *testing.T) {
	result := RemoveUncommonCharacters("\u00A1\u00A2\u00A3")
	assert.Equal(t, "", result)
}

func TestRemoveUncommonCharacters_MixedContent(t *testing.T) {
	result := RemoveUncommonCharacters("abc\u00A1def\u00A2ghi")
	assert.Equal(t, "abcdefghi", result)
}
