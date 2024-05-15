package main

// true if x = y and false otherwise

func eq() string {
	return "@sp\n" + //stack pointer.
		"A=M\n" + //get the value of the top of the stack
		"D=M\n" + // reg(D) = the value stored at the top of the stack
		"A=A-1\n" + //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
		"D=M-D\n" + //might want to do M+D if allowed because x-y x is pushed first
		"M=D\n" //load the value returned back to top of the stack
}

func gt() string {
	return "@sp\n" + //stack pointer.
		"A=M\n" + //get the value of the top of the stack.
		"D=M\n" + // reg(D) = the value stored at the top of the stack
		"A=A-1\n" + //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
		"D=M-D\n" + //might want to do M+D if allowed because x-y x is pushed first
		"M=D\n" //load the value returned back to top of the stack
}
func ls() string {
	return "@sp\n" + //stack pointer.
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
