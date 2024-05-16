package CodeWriter

import (
	"strconv"
)

// Group 1 Memory Segments: local, argument, this, that

// PopGroup1 pushes the value at the top of the stack to the specified segment offset by index
func PopGroup1(index int, segment string) string {
	return "@SP\n" + // A = 0, the location of SP
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"@" + segment + "\n" + // A = 1, the location of the segment
		"A=A+" + strconv.Itoa(index) + "\n" + // A = segment + index
		"M=D\n" // M[A] = D (the value at the top of the stack)
}

// PushGroup1 pushes the value at the specified segment offset by index to the top of the stack
func PushGroup1(index int, segment string) string {
	return "@" + segment + "\n" + // A = 1, the location of the segment
		"A=A+" + strconv.Itoa(index) + "\n" + // A = segment + index
		"D=M\n" + // D = M[A] (the value at the local segment)
		"@SP\n" + // A = 0, the top of the stack
		"A=M\n" + // A = SP (Save this pointer in A)
		"M=D\n" + // M[A] = D (the value at the local segment)
		"@SP\n" + // A = 0, the location of SP
		"M=M+1\n" // SP++ (Move stack pointer one up)
}
