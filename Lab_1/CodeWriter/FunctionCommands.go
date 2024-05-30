package CodeWriter

import (
	"fmt"
	"strconv"
)

func Function(funcName string, label string) string {
	str_to_add := "// Now in Function " + funcName + "\n" +
		"(" + funcName + ")\n"

	labelVar, err := strconv.Atoi(label)
	if err != nil {
		fmt.Println("There was a problem converting a number from string to int in the function Function of Codewriter")
	}
	for i := 0; i < labelVar; i++ {
		str_to_add += PushConstant("0")
	}
	return str_to_add
}

func Call(funcName string, argNum string, counter int) string {
	return "// Now in Call " + funcName + "\n" +
		"@RETURN_ADDRESS" + strconv.Itoa(counter) + "\n" +
		"D=A\n" + // Save return address in D
		"@SP\n" +
		"A=M\n" +
		"M=D\n" + // Push return address onto the stack
		"@SP\n" +
		"M=M+1\n" + // SP++

		PushSegmentPointer("LCL") +
		PushSegmentPointer("ARG") +
		PushSegmentPointer("THIS") +
		PushSegmentPointer("THAT") +

		"@SP\n" + // ARG = SP - 5 - argNum
		"D=M\n" +
		"@5\n" +
		"D=D-A\n" +
		"@" + argNum + "\n" +
		"D=D-A\n" +
		"@ARG\n" +
		"M=D\n" +

		"@SP\n" + // LCL = SP
		"D=M\n" +
		"@LCL\n" +
		"M=D\n" +

		"@" + funcName + "\n" + //label to jump to
		"0;JMP\n" + //jump uncondionally + // goto f

		"(RETURN_ADDRESS" + strconv.Itoa(counter) + ")\n" // (return-address)
}

func Return() string {
}
