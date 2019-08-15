package main

import (
	"fmt"

	"example.com/me/lib"
)

func main() {
	fmt.Println("")

	t := lib.MyStruct{}
	fmt.Println(t.MethodP2R2(3, 4))
}
