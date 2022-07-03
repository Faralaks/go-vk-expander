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

func TestExcludeFilenames(t *testing.T) {
	files := MsgFiles{"bob", "bob", "rob", "robert", "bert"}
	res := ExcludeFilenames(files, []string{"bob", "rob"})
	assert.EqualValues(t, MsgFiles{"robert", "bert"}, res)

	res = ExcludeFilenames(files, []string{"homer"})
	assert.EqualValues(t, MsgFiles{"bob", "bob", "rob", "robert", "bert"}, res)
}

func TestGetNumFromMsgFilename(t *testing.T) {
	// Good Way
	res, err := GetNumFromMsgFilename("messages321.html")
	assert.Equal(t, 321, res)
	assert.Nil(t, err)

	// Bad Way
	res, err = GetNumFromMsgFilename("")
	assert.NotNil(t, err)
}

func TestSortByNumber(t *testing.T) {
	files := MsgFiles{"messages50.html", "messages0.html", "messages100.html"}
	res := SortByNumber(files)
	assert.EqualValues(t, MsgFiles{"messages0.html", "messages50.html", "messages100.html"}, res)

}
