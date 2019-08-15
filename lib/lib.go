package lib

import "fmt"

// MyStruct is like the package you want to test
type MyStruct struct {
}

// MethodP0 is a member function with 0 parameter
func (m *MyStruct) MethodP0() int {
	return 1
}

// MethodP1 is a mamber function with 1 parameter
func (m *MyStruct) MethodP1(p1 int) int {
	return p1 + 100
}

// MethodP2 is a mamber function with 2 parameters
func (m *MyStruct) MethodP2(p1 int, p2 int) int {
	return p1 * p2
}

// MethodP2R2 is a member function with 2 parameters and 2 returns
func (m *MyStruct) MethodP2R2(p1 int, p2 int) (int, string) {
	return p1 * p2, fmt.Sprint(p1 * p2)
}
