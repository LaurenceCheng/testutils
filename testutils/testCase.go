package testutils

import "reflect"

// TestCase defines a test case
type TestCase struct {
	Name       string
	TargetType reflect.Type
	TestMethod string
	Args       []reflect.Value
	Want       []reflect.Value
}
