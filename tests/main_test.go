package tests

//https://translate.google.cn/  谷歌翻译
//https://write.youdao.com/edit/#/index?from=index_top  翻译
import (
	"fmt"
	"github.com/1107012776/Rookie-Learning-English/dictionary"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_MyNameIsLys(t *testing.T) {
	var obj dictionary.Introduce
	obj.MyName("YuShan Lin.")
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
	obj.Morning().I().Saw("a boy").Reading("a book.") // In the morning I saw a boy reading a book.
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_ISeeNoon(t *testing.T) {
	var obj dictionary.See
	obj.Noon().I().Saw("a boy").Reading("a book.") // At noon I saw a boy reading a book.
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
	obj.Reset()
	obj.I().Saw("a boy").Reading("a book").Noon().Dot() // I saw a boy reading a book at noon.
	str, err = obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}

func Test_I_Am(t *testing.T) {
	var obj dictionary.Introduce
	obj.I_Am("a developer")
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)

}

func Test_YouShould(t *testing.T) {
	var obj dictionary.See
	obj.YouShould("a developer").Dot() //You should see a developer. 你应该看到一个开发者
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
	obj.Reset()
	obj.YouShould("something like that").Dot() //You should see something like that.  你应该看到类似的东西
	str, err = obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)

}

func Test_AddConfigurationData(t *testing.T) {
	var obj dictionary.Add
	obj.AddConfigurationData("php.ini").Dot() //Add configuration data on php.ini. 在 php.ini 上添加配置数据
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}
