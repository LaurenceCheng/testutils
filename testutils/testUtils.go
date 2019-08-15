package testutils

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// DeepEqual is a method
func DeepEqual(actual []reflect.Value, expect []reflect.Value) bool {
	if actual == nil || expect == nil {
		return reflect.DeepEqual(actual, expect)
	}

	actualLen := len(actual)
	expectLen := len(expect)
	if actualLen != expectLen {
		return false
	}

	for i := 0; i < actualLen; i++ {
		if !reflect.DeepEqual(actual[i].Interface(), expect[i].Interface()) {
			return false
		}
	}

	return true
}

// RunTests runs all tests the user defined
func RunTests(t *testing.T, tests []TestCase) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := reflect.New(tt.TargetType).MethodByName(tt.TestMethod).Call(tt.Args); !DeepEqual(got, tt.Want) {
				t.Errorf("%v.%s() = %v, want %v", tt.TargetType, tt.TestMethod, toString(got), toString(tt.Want))
			}
		})
	}

}

// GetType ...
func GetType(t interface{}) reflect.Type {
	return reflect.TypeOf(t)
}

// GetWant ...
func GetWant(wants ...interface{}) []reflect.Value {
	return getArrayOfReflectValue(wants)
}

// GetArgs ...
func GetArgs(args ...interface{}) []reflect.Value {
	return getArrayOfReflectValue(args)
}

func getArrayOfReflectValue(items []interface{}) []reflect.Value {
	result := make([]reflect.Value, 0)

	for _, i := range items {
		result = append(result, reflect.ValueOf(i))
	}
	return result
}

func toString(values []reflect.Value) string {
	builder := strings.Builder{}
	builder.Write([]byte("["))

	strArray := make([]string, 0)
	for _, v := range values {
		toString := fmt.Sprintf("%v (%s) ", v.Interface(), v.Kind())
		strArray = append(strArray, toString)
	}

	builder.WriteString(strings.Join(strArray, ", "))

	builder.Write([]byte("]"))
	return builder.String()
}
