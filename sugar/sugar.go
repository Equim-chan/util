// Package sugar provides some syntactic-sugar-like utilities
package sugar

// Must panics if the error is non-nil. It is intended for use in variable
// initializations and assertions.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Same as Must but accept two parameters.
// Example:
//     literal := Must2(syscall.UTF16PtrFromString("Literal")).(*uint16)
func Must2(f interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return f
}
