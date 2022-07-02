package html_extractor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNameInList(t *testing.T) {
	res := IsNameInList("robert", []string{"bob", "rob", "robert", "bert"})
	assert.True(t, res)

	res = IsNameInList("homer", []string{"bob", "rob", "robert", "bert"})
	assert.False(t, res)
}

func TestFiles_ExcludeFilenames(t *testing.T) {
	files := Files{"bob", "bob", "rob", "robert", "bert"}
	res := files.ExcludeFilenames([]string{"bob", "rob"})
	assert.EqualValues(t, Files{"robert", "bert"}, res)

	res = files.ExcludeFilenames([]string{"homer"})
	assert.EqualValues(t, Files{"bob", "bob", "rob", "robert", "bert"}, res)
}

func TestGetNumFromMsgFilename(t *testing.T) {
	// Good Way
	res, err := GetNumFromMsgFilename("message321.html")
	assert.Equal(t, 321, res)
	assert.Nil(t, err)

	// Bad Way
	res, err = GetNumFromMsgFilename("")
	assert.NotNil(t, err)
}
