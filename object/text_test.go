package object_test

import (
	"github.com/louvri/gob/object"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomText(t *testing.T) {
	text1 := object.GenerateRandomText(4)
	text2 := object.GenerateRandomText(4)
	text3 := object.GenerateRandomText(4)
	text4 := object.GenerateRandomText(4)
	assert.NotEqual(t, text1, text2)
	assert.NotEqual(t, text2, text3)
	assert.NotEqual(t, text3, text4)
	assert.NotEqual(t, text1, text3)
	assert.NotEqual(t, text1, text4)
	assert.NotEqual(t, text2, text4)
}
