package main

// true if x = y and false otherwise

func eq() string {
	var str string = ""
	str += "@sp\n"   //stack pointer.
	str += "A=M\n"   //get the value of the top of the stack
	str += "D=M\n"   // reg(D) = the value stored at the top of the stack
	str += "A=A-1\n" //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
	str += "D=M-D\n" //might want to do M+D if allowed because x-y x is pushed first
	str += "M=D\n"   //load the value returned back to top of the stack
	return str
}

//realized that eq implements just equal this is after a push  x and push y happened
// func eq(loc_x string, loc_y string, loc_jump string)string {
// 	return "@" + loc_x + "\nD = M\n@" + loc_y + "D = D - M\n"
// }
