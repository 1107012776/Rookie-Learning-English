package tests

import (
	"fmt"
	"github.com/1107012776/Rookie-Learning-English/dictionary"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_MyNameIsLys(t *testing.T) {
	var introduce dictionary.Introduce
	str, err := introduce.Yourself("YuShanLin")
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}
