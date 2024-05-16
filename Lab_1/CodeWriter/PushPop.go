package CodeWriter

import (
	"strconv"
)

// PopLocal pushes the value at the top of the stack to the local segment offset by index
func PopLocal(index int) string {
	return "@SP\n" + // A = 0, the location of SP
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"@LCL\n" + // A = 1, the location of LCL
		"A=A+" + strconv.Itoa(index) + "\n" + // A = LCL + index
		"M=D\n" // M[A] = D (the value at the top of the stack)
}

// PushLocal pushes the value at the local segment offset by index to the top of the stack
func PushLocal(index int) string {
	return "@LCL\n" + // A = 1, the location of LCL
		"A=A+" + strconv.Itoa(index) + "\n" + // A = LCL + index
		"D=M\n" + // D = M[A] (the value at the local segment)
		"@SP\n" + // A = 0, the top of the stack
		"A=M\n" + // A = SP (Save this pointer in A)
		"M=D\n" + // M[A] = D (the value at the local segment)
		"@SP\n" + // A = 0, the location of SP
		"M=M+1\n" // SP++ (Move stack pointer one up)
}
