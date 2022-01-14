package dictionary

/**
添加结构体
*/
type Add struct {
	Base
}

/**
在 “什么” 上添加配置数据
*/
func (obj *Add) AddConfigurationData(name string) *Add {
	obj.content = "Add configuration data on " + name
	return obj
}
