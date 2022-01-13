package dictionary

/**
介绍结构体
*/
type Introduce struct {
}

func (obj *Introduce) A_Person(name string) {

}

/**
介绍自己
*/
func (obj *Introduce) Yourself(name string) (string, error) {
	return "My name is " + name, nil
}
