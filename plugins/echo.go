package plugins

import (
	"reflect"
)

type EchoPlugin struct {
	mapping map[int]interface{}
}

func (plugin *EchoPlugin) Init(mapping map[int]interface{}) {
	mapping[0] = reflect.String
}

func (plugin *EchoPlugin) Mapping() map[int]interface{} {
	if plugin.mapping == nil {
		plugin.mapping = make(map[int]interface{})
	}
	return plugin.mapping
}

func (EchoPlugin) Name() string {
	return "echo"
}

func (plugin EchoPlugin) Invoke(args ...interface{}) string {
	if len(args) == 1 {
		return " (动态输出插件) -> " + args[0].(string)
	}
	return ""
}
