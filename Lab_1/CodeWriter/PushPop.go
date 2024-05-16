package CodeWriter

import (
	"strconv"
)

// Group 1 Memory Segments: local, argument, this, that

func PushArgument(index int) string {
	return pushGroup1(index, "ARG")
}

func PopArgument(index int) string {
	return popGroup1(index, "ARG")
}

func PushLocal(index int) string {
	return pushGroup1(index, "LCL")
}

func PopLocal(index int) string {
	return popGroup1(index, "LCL")
}

func PushThis(index int) string {
	return pushGroup1(index, "THIS")
}

func PopThis(index int) string {
	return popGroup1(index, "THIS")
}

func PushThat(index int) string {
	return pushGroup1(index, "THAT")
}

func PopThat(index int) string {
	return popGroup1(index, "THAT")
}

// PopGroup1 pop the top of the stack into the segment offset by index
func popGroup1(index int, segment string) string {
	return "@SP\n" + // A = 0, the location of SP
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"@" + segment + "\n" + // A = the location of the segment
		"A=A+" + strconv.Itoa(index) + "\n" + // A = segment + index
		"M=D\n" // M[A] = D (the value at the top of the stack)
}

// PushGroup1 push the value at the specified segment offset by index to the top of the stack
func pushGroup1(index int, segment string) string {
	return "@" + segment + "\n" + // A = the location of the segment
		"A=A+" + strconv.Itoa(index) + "\n" + // A = segment + index
		"D=M\n" + // D = M[A] (the value at the segment)
		"@SP\n" + // A = 0, the top of the stack
		"A=M\n" + // A = SP (Save this pointer in A)
		"M=D\n" + // M[A] = D (the value at the segment)
		"@SP\n" + // A = 0, the location of SP
		"M=M+1\n" // SP++ (Move stack pointer one up)
}

// Group 2: temp

// pop the top of the stack into address RAM[ 5 + x ]
func PopTemp(index int) string {
	return "@SP\n" + // A = 0, the location of SP
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"@" + strconv.Itoa(5+index) + "\n" + // A = 5 + x
		"M=D\n" // M[A] = D (the value at the top of the stack)
}

// push the value at address RAM[ 5 + x ] onto the stack
func PushTemp(index int) string {
	return "@" + strconv.Itoa(5+index) + "\n" + // A = 5 + x
		"D=M\n" + // D = M[A] (the value at the address 5 + x)
		"@SP\n" + // A = 0, the top of the stack
		"A=M\n" + // A = SP (Save this pointer in A)
		"M=D\n" + // M[A] = D (the value at the address 5 + x)
		"@SP\n" + // A = 0, the location of SP
		"M=M+1\n" // SP++ (Move stack pointer one up)
}

// Group 3: static

// pop the top of the stack into address file_name.index
func PopStatic(index int, file_name string) string {
	return "@SP\n" + // A = 0, the location of SP
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"@" + file_name + "." + strconv.Itoa(index) + "\n" + // A = file_name.index
		"M=D\n" // M[A] = D (the value at the top of the stack)
}

// push the value at address file_name.index onto the stack
func PushStatic(index int, file_name string) string {
	return "@" + file_name + "." + strconv.Itoa(index) + "\n" + // A = file_name.index
		"D=M\n" + // D = M[A] (the value at the address file_name.index)
		"@SP\n" + // A = 0, the top of the stack
		"A=M\n" + // A = SP (Save this pointer in A)
		"M=D\n" + // M[A] = D (the value at the address file_name.inde)
		"@SP\n" + // A = 0, the location of SP
		"M=M+1\n" // SP++ (Move stack pointer one up)
}

// Group 4: pointer

// pop the top of the stack into address RAM[THIS/THAT]
// index = 0: THIS, index = 1: THAT
func PopPointer(index int) string {
	segment := "THIS"
	if index == 1 {
		segment = "THAT"
	}

	return "@SP\n" + // A = 0, the location of SP
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"@" + segment + "\n" + // A = THIS/THAT
		"M=D\n" // M[A] = D (the value at the top of the stack)
}

// push the value at address THIS/THAT onto the stack
// index = 0: THIS, index = 1: THAT
func PushPointer(index int) string {
	segment := "THIS"
	if index == 1 {
		segment = "THAT"
	}

	return "@" + segment + "\n" + // A = THIS/THAT
		"D=M\n" + // D = M[A] (the value at THIS/THAT)
		"@SP\n" + // A = 0, the top of the stack
		"A=M\n" + // A = SP (Save this pointer in A)
		"M=D\n" + // M[A] = D (the value at THIS/THAT)
		"@SP\n" + // A = 0, the location of SP
		"M=M+1\n" // SP++ (Move stack pointer one up)
}
