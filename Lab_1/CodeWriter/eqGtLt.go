package CodeWriter

// true if x = y and false otherwise

func Eq() string {
	return "@SP\n" + //stack pointer.
		"A=M\n" + //get the value of the top of the stack
		"D=M\n" + // reg(D) = the value stored at the top of the stack
		"A=A-1\n" + //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
		"D=M-D\n" + //might want to do M+D if allowed because x-y x is pushed first
		"M=D\n" //load the value returned back to top of the stack
}

// Attempt at Eq by Avi
func Eq() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
		"D=M-D\n" + // D = M[A] - D (the value at the top of the stack)
		"@Equal\n" + // jump to Equal if x = y (x-y = 0)
		"D;JEQ\n" + // if x != y, jump to Equal
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
		"M=-1\n" + // M[A] = -1 top of the stack is -1 as not equal
		"@End\n" + // jump to End
		"0;JMP\n" + // jump to End
		"(Equal)\n" + // label Equal
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
		"M=0\n" + // M[A] = 0 top of the stack is 0 as equal
		"(End)\n" // label End
}

func Gt() string {
	return "@SP\n" + //stack pointer.
		"A=M\n" + //get the value of the top of the stack.
		"D=M\n" + // reg(D) = the value stored at the top of the stack
		"A=A-1\n" + //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
		"D=M-D\n" + //might want to do M+D if allowed because x-y x is pushed first
		"M=D\n" //load the value returned back to top of the stack
}
func Lt() string {
	return "@SP\n" + //stack pointer.
		"A=M\n" + //get the value of the top of the stack
		"D=M\n" + // reg(D) = the value stored at the top of the stack
		"A=A-1\n" + //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
		"D=M-D\n" + //might want to do M+D if allowed because x-y x is pushed first
		"M=D\n" //load the value returned back to top of the stack
}

//realized that eq implements just equal this is after a push  x and push y happened
// func eq(loc_x string, loc_y string, loc_jump string)string {
// 	return "@" + loc_x + "\nD = M\n@" + loc_y + "D = D - M\n"
// }
