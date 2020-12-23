package config

import (
	"github.com/4everlynn/journal/plugins"
)

// Runtime structure
type Runtime struct {
	Plugins map[string]plugins.PluginFunc
}

// runtime singleton instance
var runtime = Runtime{
	Plugins: map[string]plugins.PluginFunc{},
}

// GetRuntime runtime instance
func GetRuntime() Runtime {
	return runtime
}

// Install install plugin
func Install(pluginFunc plugins.PluginFunc) {
	pluginFunc.Init(pluginFunc.Mapping())
	GetRuntime().Plugins[pluginFunc.Name()] = pluginFunc
}

// Invoke trigger plugin function
func Invoke(name string, args ...interface{}) string {
	pluginFunc := GetRuntime().Plugins[name]
	if pluginFunc != nil {
		return pluginFunc.Invoke(args...)
	}
	return ""
}

// InstallDefaultPlugins install the default plugin
func InstallDefaultPlugins() {
	Install(&plugins.EchoPlugin{})
}
