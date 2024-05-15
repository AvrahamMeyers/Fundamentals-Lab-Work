package main

//func add() string {
//	return "@sp\n" +
//		"D = M\n" +
//		"A = A - 1\n" +
//		"MD = D + M\n"
//}

func add() string {
	return "@sp\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the second top of the stack)
		"@sp\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
		"M=D+M\n" // M[A] = D + M[A] (the value at the top of the stack)
}

//func sub() string {
//	return "@x \n" +
//		"D = A\n" +
//		"@y \n" +
//		"D = D - A\n"
//}

func sub() string {
	return "@sp\n" +
		"D = A\n" +
		"@y \n" +
		"D = D - A\n"
}

func neg() string {
	return "@y \n" +
		"D = -A\n"
}
