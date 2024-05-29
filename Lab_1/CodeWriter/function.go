package CodeWriter

import (
	"fmt"
	"strconv"
)

func Function(fname string, lbl string) string {
	str_to_add := "// Now in Function " + fname + "\n(" + fname + ")\n"
	lclVar, err := strconv.Atoi(lbl)
	if err != nil {
		fmt.Println("There was a problem converting a number from string to int in the function Function of Codewriter")
	}
	for i := 0; i < lclVar; i++ {
		str_to_add += PushConstant("0")
	}
	return str_to_add
}
