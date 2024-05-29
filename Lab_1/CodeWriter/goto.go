package CodeWriter

func Goto(fnction string, lbl string) string {
	return "@" + fnction + "$" + lbl + "\n" + //lbl to jump to
		"D;JMP\n" //jump uncondionally

}

func GotoIf(fnction string, lbl string) string {
	return "D=D+1\n" +
		"@" + fnction + "$" + lbl + "\n" +
		"D;JEQ\n"
	//may need to pop one of the stack. will find out during testing.
}
