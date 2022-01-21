# Rookie-Learning-English
菜鸟英语学习工具，将英文语法进行封装，通过代码来写英文
Rookie English learning tool, encapsulates English grammar and writes English through code.

# Example
```golang
package tests


import (
	"fmt"
	"github.com/1107012776/Rookie-Learning-English/dictionary"
	assert "github.com/magiconair/properties/assert"
	"testing"
)
func Test_I_Am(t *testing.T) {
	var obj dictionary.Introduce
	obj.I_Am("a developer").Dot()  //I am a developer.
	str, err := obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
	obj.Reset()
	obj.MyName("Yushan Lin").Dot()  //My name is Yushan Lin.
	str, err = obj.Response()
	fmt.Println(str)
	assert.Equal(t, err == nil, true)
}
```
