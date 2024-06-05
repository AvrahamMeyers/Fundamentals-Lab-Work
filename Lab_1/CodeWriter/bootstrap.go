package CodeWriter

func Bootsrap() string {

	return "@256\n" +
		"D=A\n" +
		"@SP\n" +
		"M=D\n" +
		"// Now in Call " + "Sys.init" + "\n" +
		"@RETURN_ADDRESS" + "0" + "\n" +
		"D=A\n" + // Save return address in D
		"@SP\n" +
		"A=M\n" +
		"M=D\n" + // Push return address onto the stack
		"@SP\n" +
		"M=M+1\n" + // SP++

		PushSegmentPointer("LCL") +
		PushSegmentPointer("ARG") +
		PushSegmentPointer("THIS") +
		PushSegmentPointer("THAT") +

		"@SP\n" + // ARG = SP - 5 - argNum
		"D=M\n" +
		"@5\n" +
		"D=D-A\n" +
		"@" + "0" + "\n" +
		"D=D-A\n" +
		"@ARG\n" +
		"M=D\n" +

		"@SP\n" + // LCL = SP
		"D=M\n" +
		"@LCL\n" +
		"M=D\n" +

		"@" + "Sys.init" + "\n" + //label to jump to
		"0;JMP\n" + //jump uncondionally + // goto f

		"(RETURN_ADDRESS" + "0" + ")\n" // (return-address)
	//"@Sys.init\n" +
	//"D;JMP\n"

}
