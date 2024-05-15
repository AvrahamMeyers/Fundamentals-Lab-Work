package main

func gt() string {
	var str string = ""
	str += "@sp\n"   //stack pointer
	str += "A=M\n"   //get the value of the top of the stack
	str += "D=M\n"   // reg(D) = the value stored at the top of the stack
	str += "A=A-1\n" //move the stack down one or perhaps it should be the opposite up one depends on stack implementation
	str += "D=D-M\n" //might want to do M+D if allowed because x-y x is pushed first
	str += "M=D\n"   //load the value returned back to top of the stack
	return str
}
