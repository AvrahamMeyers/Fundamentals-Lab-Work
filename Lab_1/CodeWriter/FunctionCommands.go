package CodeWriter

import (
	"fmt"
	"strconv"
)

func Function(funcName string, label string) string {
	str_to_add := "// Now in Function " + funcName + "\n(" + funcName + ")\n"
	lclVar, err := strconv.Atoi(label)
	if err != nil {
		fmt.Println("There was a problem converting a number from string to int in the function Function of Codewriter")
	}
	for i := 0; i < lclVar; i++ {
		str_to_add += PushConstant("0")
	}
	return str_to_add
}

func Call(funcName string, argNum string, counter int) string {
	return "// Now in Call " + funcName + "\n" +
		"@RETURN_ADDRESS" + strconv.Itoa(counter) + "\n" +
		"D=A\n" +
		"@SP\n" +
		"A=M\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M+1\n" +
		PushConstant("LCL") +
		PushConstant("ARG") +
		PushConstant("THIS") +
		PushConstant("THAT") +
		"@SP\n" +
		"D=M\n" +
		"@5\n" +
		"D=D-A\n" +
		"@" + strconv.Itoa(argNum) + "\n" +
		"D=D-A\n" +
		"@ARG\n" +
		"M=D\n" +
		"@SP\n" +
		"D=M\n" +
		"@LCL\n" +
		"M=D\n" +
		Goto(funcName) +
		"(RETURN_ADDRESS" + strconv.Itoa(counter) + ")\n"
}

func Return() string {
}
