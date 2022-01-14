package dictionary

/**
介绍结构体
*/
type Introduce struct {
	Base
}

func (obj *Introduce) A_Person(name string) {

}

func (obj *Introduce) MyName(name string) *Introduce {
	obj.content = "My name is " + name
	return obj
}

func (obj *Introduce) I_Am(str string) *Introduce {
	obj.content = "I am " + str
	return obj
}
