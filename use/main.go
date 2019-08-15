package main

import (
	"fmt"

	"github.com/LaurenceCheng/testutils/lib"
)

func main() {
	fmt.Println("")

	t := lib.MyStruct{}
	fmt.Println(t.MethodP2R2(3, 4))
}
