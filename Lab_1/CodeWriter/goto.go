package CodeWriter

// funcName is also the scope
func Goto(funcName string, label string) string {
	return "@" + funcName + "$" + label + "\n" + //label to jump to
		"D;JMP\n" //jump uncondionally
}

// funcName is also the scope
func GotoIf(funcName string, label string) string {
	return "D=D+1\n" +
		"@" + funcName + "$" + label + "\n" +
		"D;JEQ\n"
	//may need to pop one of the stack. will find out during testing.
}
