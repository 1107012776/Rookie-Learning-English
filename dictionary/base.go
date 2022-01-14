package dictionary

type Base struct {
	content string
	err     error
}

func (obj *Base) Reset() *Base {
	obj.content = ""
	obj.err = nil
	return obj
}

func (obj *Base) Response() (string, error) {
	return obj.content, obj.err
}

/**
点号
*/
func (obj *Base) Dot() *Base {
	obj.content += "."
	return obj
}

/**
逗号
*/
func (obj *Base) Comma() *Base {
	obj.content += ","
	return obj
}
