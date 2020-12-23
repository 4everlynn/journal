package support

import (
	"errors"
	"github.com/4everlynn/journal/config"
	"reflect"
	"regexp"
	"strings"
)

// FuncParser convert string to func
type FuncParser interface {
	// msg string in format as @func(...args)
	Parse(msg string) config.Func
}

type SmartFuncParser struct{}

func (SmartFuncParser) Parse(msg string) (config.Func, error) {
	compile, err := regexp.Compile("@(\\w+)\\((.*?)\\)+")

	if err == nil {
		matched := compile.FindStringSubmatch(msg)
		if len(matched) == 3 {
			cmd := matched[1]
			mFunc := config.GetRuntime().Plugins[cmd]
			if mFunc != nil {
				params := matched[2]
				split := strings.Split(params, ",")
				mapping := mFunc.Mapping()
				args := make([]interface{}, len(mapping))
				if mapping != nil && len(mapping) > 0 && len(split) == len(mapping) {
					for index, _ := range split {
						v := split[index]
						args[index] = ParseKind(v, mapping[index].(reflect.Kind))
					}
					return config.Func{
						Name: matched[1],
						Args: args,
					}, nil
				}
			}
		}
	}

	return config.Func{}, errors.New("error when parsing plugin-func")
}
