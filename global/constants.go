package global

// Separator define a cross platform file separator
var Separator = _injectSep()

func _injectSep() string {
	switch GetRuntime() {
	case LINUX, OSX:
		return "/"
	case WINDOWS:
		return "\\"
	}
	return "/"
}
