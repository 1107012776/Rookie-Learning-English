package dictionary

//ThereBe Sentence Patterns  (ThereBe 句型)
type ThereBe struct {
	Base
}

func (obj *ThereBe) ThereAre(str string) *ThereBe {
	obj.content = "There are " + str
	return obj
}

func (obj *ThereBe) ThereIs(str string) *ThereBe {
	obj.content = "There is " + str
	return obj
}
