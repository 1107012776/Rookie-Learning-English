package tests

//https://translate.google.cn/  谷歌翻译
//https://write.youdao.com/edit/#/index?from=index_top  翻译 有道写作
import (
	"fmt"
	"github.com/1107012776/Rookie-Learning-English/dictionary"
	"github.com/1107012776/Rookie-Learning-English/dictionary/animal"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_Shark(t *testing.T) {
	var obj dictionary.ThereBe
	obj.ThereAre("some sharks").Dot() //那里有一些鲨鱼。
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_Sharks(t *testing.T) {
	var obj animal.Shark
	obj.There_are_some_sharks().Dot() //那里有一些鲨鱼。
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_Book(t *testing.T) {
	var obj dictionary.ThereBe
	obj.ThereAre("some books").Dot() //那里有一些书。
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_A_Book(t *testing.T) {
	var obj dictionary.ThereBe
	obj.ThereIs("a book").Dot() //那里有一本书。
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}
