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
