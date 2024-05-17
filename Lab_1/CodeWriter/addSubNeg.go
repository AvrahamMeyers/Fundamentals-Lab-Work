package CodeWriter

func Add() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=D+M\n" // M[A] = D + M[A] (the value at the top of the stack)
}

func Sub() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=M-D\n" // M[A] = M[A] - D (the value at the top of the stack)
}

func Neg() string {
	return "@SP\n" +
		"A=M-1\n" +
		"M=-M\n"
}
