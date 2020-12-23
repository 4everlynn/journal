package plugins

// PluginFunc plugin function
type PluginFunc interface {
	Mapping() map[int]interface{}
	Init(mapping map[int]interface{})
	Name() string
	Invoke(args ...interface{}) string
}
