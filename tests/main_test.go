package tests

//https://write.youdao.com/edit/#/index?from=index_top  翻译
import (
	"fmt"
	"github.com/1107012776/Rookie-Learning-English/dictionary"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_MyNameIsLys(t *testing.T) {
	var obj dictionary.Introduce
	obj.MyName("YuShanLin")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_ISee(t *testing.T) {
	var obj dictionary.See
	obj.I().See("a boy").Reading("a book")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_ISeeYesterday(t *testing.T) {
	var obj dictionary.See
	obj.I().Saw("a boy").Reading("a book").Yesterday("")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_ISeeMorning(t *testing.T) {
	var obj dictionary.See
	str, _ := obj.Response()
	fmt.Println(str)
	obj.Morning().I().Saw("a boy").Reading("a book")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_ISeeNoon(t *testing.T) {
	var obj dictionary.See
	obj.Noon().I().Saw("a boy").Reading("a book")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
	obj.Reset()
	obj.I().Saw("a boy").Reading("a book").Noon()
	str, err = obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_I_am(t *testing.T) {
	var obj dictionary.Introduce
	obj.I_Am("a developer")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)

}
