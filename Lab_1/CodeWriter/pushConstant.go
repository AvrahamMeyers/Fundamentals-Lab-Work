package CodeWriter

//expected is just the number string return the commands @num  D=A
func PushConstant(num string) string {
	return "@" + num + "\n" + //set A = to number
		"D=A\n" + //set D = to num
		"@SP\n" + //set A = SP (0)
		"A=M\n" + //set A = RAM[SP]
		"M=D\n" + //set RAM[Ram[SP]] = to num
		"@SP\n" + //A = SP
		"M=M+1\n" //set SP to be +1
}
