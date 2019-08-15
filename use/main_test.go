package main

import (
	"testing"

	"github.com/LaurenceCheng/testutils/lib"
	tu "github.com/LaurenceCheng/testutils/testutils"
)

func TestMethodP0(t *testing.T) {
	tests := []tu.TestCase{
		{
			Name:       "TestMyStructMethodP0",
			TargetType: tu.GetType(lib.MyStruct{}),
			TestMethod: "MethodP0",
			Want:       tu.GetWant(1),
		},
	}

	tu.RunTests(t, tests)
}

func TestMethodP1(t *testing.T) {
	tests := []tu.TestCase{
		{
			Name:       "TestMyStructMethodP1",
			TargetType: tu.GetType(lib.MyStruct{}),
			TestMethod: "MethodP1",
			Args:       tu.GetArgs(1),
			Want:       tu.GetWant(101),
		},
		{
			Name:       "TestMyStructMethodP1Minus1",
			TargetType: tu.GetType(lib.MyStruct{}),
			TestMethod: "MethodP1",
			Args:       tu.GetArgs(-1),
			Want:       tu.GetWant(100),
		},
	}

	tu.RunTests(t, tests)
}

func TestMethodP2(t *testing.T) {
	tests := []tu.TestCase{
		{
			Name:       "TestMyStructMethodP2",
			TargetType: tu.GetType(lib.MyStruct{}),
			TestMethod: "MethodP2",
			Args:       tu.GetArgs(3, 4),
			Want:       tu.GetWant(15),
		},
	}

	tu.RunTests(t, tests)
}

func TestMethodP2R2(t *testing.T) {
	tests := []tu.TestCase{
		{
			Name:       "TestMyStructMethodP2R2",
			TargetType: tu.GetType(lib.MyStruct{}),
			TestMethod: "MethodP2R2",
			Args:       tu.GetArgs(3, 4),
			Want:       tu.GetWant(12, 12),
		},
	}

	tu.RunTests(t, tests)
}