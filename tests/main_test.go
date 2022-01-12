package tests

import (
	"fmt"
	"github.com/1107012776/rookie-learning-english/dictionary"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_MyNameIsLys(t *testing.T) {
	str, err := dictionary.Introduce_Yourself("YuShanLin")
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}
