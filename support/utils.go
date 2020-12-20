package support

// Catch global catch func
func Catch(err error, hook func()) {
	if err != nil {
		panic(err)
	}
	if hook != nil {
		hook()
	}
}
