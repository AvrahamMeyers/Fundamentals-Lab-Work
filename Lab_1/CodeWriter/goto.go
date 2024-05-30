package CodeWriter

func Goto(funcName string, label string) string {
	return "@" + funcName + "$" + label + "\n" + //lbl to jump to
		"D;JMP\n" //jump uncondionally

}

func GotoIf(funcName string, label string) string {
	return "D=D+1\n" +
		"@" + funcName + "$" + label + "\n" +
		"D;JEQ\n"
	//may need to pop one of the stack. will find out during testing.
}
