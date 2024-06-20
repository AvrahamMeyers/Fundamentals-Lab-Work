package CodeWriter

// funcName is also the scope
func Goto(funcName string, label string) string {
	return "@" + funcName + "$" + label + "\n" + //label to jump to
		"D;JMP\n" //jump uncondionally
}

// funcName is also the scope
func GotoIf(funcName string, label string) string {
	return "@SP\n" +
		"A=M\n" + //A = SP address
		"A=A-1\n" + //A = SP - 1
		"D=M\n" + //D = value at top of stack
		"@SP\n" +
		"M=M-1\n" + //SP--
		"@" + funcName + "$" + label + "\n" +
		"D;JNE\n" // jump if D != 0
}
