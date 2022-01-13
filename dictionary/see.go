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
	if obj.content == "" {
		obj.content = " I"
	} else {
		obj.content += " I"
	}
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

func (obj *See) Morning() *See {
	if obj.content == "" {
		obj.content = " In the morning"
		return obj
	}
	obj.content += " this morning"
	return obj
}

func (obj *See) Noon() *See {
	if obj.content == "" {
		obj.content = " At noon"
		return obj
	}
	obj.content += " at noon"
	return obj
}
