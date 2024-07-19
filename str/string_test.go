package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractAlfaNumericWithSelectedSpecialCharactersFromText(t *testing.T) {
	sourceText := "Ab9 \\!@#$%^&*_+-=\n"
	validateText := "Ab9 \\!#$%&*+-\n"
	filteredText := ExtractAlfaNumericWithSelectedSpecialCharactersFromText(sourceText)
	assert.Equal(t, validateText, filteredText)
}
