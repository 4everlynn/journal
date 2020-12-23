package support

import (
	"reflect"
	"strconv"
	"strings"
)

// ParseKind
func ParseKind(str string, kind reflect.Kind) interface{} {
	switch kind {
	case reflect.Int:
		v, _ := strconv.Atoi(str)
		return v
	case reflect.Float64:
		v, _ := strconv.ParseFloat(str, 64)
		return v
	case reflect.Bool:
		v, _ := strconv.ParseBool(str)
		return v
	default:
		return strings.Trim(str, " ")
	}
}
