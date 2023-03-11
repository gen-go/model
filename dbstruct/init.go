package dbstruct

var structMap = make(map[string]interface{})

// Get 根据结构体名获取对象
func Get(name string) interface{} {
	return structMap[name]
}

func set(name string, obj interface{}) {
	structMap[name] = obj
}
