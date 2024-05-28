package CodeWriter

func Bootsrap() string {
	return "@256\n" +
		"D=A\n" +
		"@SP\n" +
		"M=D\n" +
		"@Sys.init\n" +
		"D;jmp\n"

}
