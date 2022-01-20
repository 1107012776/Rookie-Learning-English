package animal

import "github.com/1107012776/Rookie-Learning-English/dictionary"

//鲨鱼
type Shark struct {
	dictionary.Base
	name string
}

func (obj *Shark) There_are_some_sharks() *Shark {
	obj.SetContent("There are some sharks") //这有鲨鱼
	return obj
}
