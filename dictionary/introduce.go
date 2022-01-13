package dictionary

/**
介绍结构体
*/
type Introduce struct {
	content string
	err     error
}

func (obj *Introduce) A_Person(name string) {

}

func (obj *Introduce) MyName(name string) *Introduce {
	obj.content = "My name is " + name
	return obj
}

func (obj *Introduce) Response() (string, error) {
	return obj.content, obj.err
}
