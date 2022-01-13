package dictionary

/**
看
*/

/**
看结构体
*/
type See struct {
	content string
	err     error
}

func (obj *See) I() *See {
	obj.content = "I"
	return obj
}

func (obj *See) Saw(name string) *See {
	obj.content += " saw " + name
	return obj
}

func (obj *See) See(name string) *See {
	obj.content += " see " + name
	return obj
}

func (obj *See) Reading(name string) *See {
	obj.content += " reading " + name
	return obj
}

func (obj *See) Yesterday(str string) *See {
	obj.content += " yesterday " + str
	return obj
}

func (obj *See) Response() (string, error) {
	return obj.content, obj.err
}
