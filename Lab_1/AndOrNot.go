package main

func and() string {
	return "@sp\n" + // A = 0, the location of SP
		"A=M-1\n" + // SP-- (Move stack pointer one down)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@sp\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=D&M\n" // M[A] = D and M[A] (the value at the top of the stack)
	// + "@sp\n" + // increment the stack pointer to the empty slot above the result of the operation
	// "M=M+1\n"
}

func or() string {
	return "@sp\n" + // A = 0, the location of SP
		"A=M-1\n" + // SP-- (Move stack pointer one down)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@sp\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=D|M\n" // M[A] = D and M[A] (the value at the top of the stack)
	// + "@sp\n" + // increment the stack pointer to the empty slot above the result of the operation
	// "M=M+1\n"
}

func not() string {
	return "@sp\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
		"M=!M\n" // M[A] = not M[A] (the value at the top of the stack)
	// + "@sp\n" + // increment the stack pointer to the empty slot above the result of the operation
	// "M=M+1\n"
}
