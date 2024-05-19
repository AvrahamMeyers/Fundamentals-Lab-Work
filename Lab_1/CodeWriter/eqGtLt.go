package CodeWriter

import "strconv"

func Eq(counter int) string {
	return "@SP\n" + // A=0
		"M=M-1\n" + //point sp at top
		"A=M\n" + //set mem to sp
		"D=M\n" + // SAVE THE TOP OF THE STACK
		"A=A-1\n" + //POP MOVE A TO THE NEXT ITEM IN STACK
		"D=M-D\n" + //X-Y FOR THE 2 STACK NUMS
		"@IF_TRUE" + strconv.Itoa(counter) + "\n" + //jump label
		"D;JEQ\n" + // JUMP IF D = 0
		"D=0\n" + //if not equal set d=0
		"@END" + strconv.Itoa(counter) + "\n" + //SET SUMP LOACTION
		"D;JMP\n" + //JUMP TO THE END
		"(IF_TRUE" + strconv.Itoa(counter) + ")\n" + // true label
		"D=-1\n" + //SET D=-1 THOUGHT IT WOULD BE 0 BUT THE TESTS WANT -1
		"(END" + strconv.Itoa(counter) + ")" + "\n" + //End label
		"@SP\n" + //A=0
		"A=M\n" + //set A of Ram[A] set to point at top of stack (+1)
		"A=A-1\n" + // Now point at the topmost elelment of the stack which was last element popped
		"M=D\n" // set the top of the stack to appropriate logical value
}

//Attempt at Eq by Avi
// func Eq() string {
// 	return "@SP\n" + // A = 0, the location of SP
// 		"M=M-1\n" + // SP-- (Move stack pointer one down)
// 		"A=M\n" + // A = SP (Save this pointer in A)
// 		"D=M\n" + // D = M[A] (the value at the top of the stack)
// 		"@SP\n" + // A = 0, the location of SP
// 		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
// 		"D=M-D\n" + // D = M[A] - D (the value at the top of the stack)
// 		"@Equal\n" + // jump to Equal if x = y (x-y = 0)
// 		"D;JEQ\n" + // if x != y, jump to Equal
// 		"@SP\n" + // A = 0, the location of SP
// 		"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
// 		"M=-1\n" + // M[A] = -1 top of the stack is -1 as not equal
// 		"@End\n" + // jump to End
// 		"0;JMP\n" + // jump to End
// 		"(Equal)\n" + // label Equal
// 		// should not move sp again its one too many times"@SP\n" + // A = 0, the location of SP
// 		//"A=M-1\n" + // A = M[A] - 1 (the location of the top of the stack)
// 		"M=0\n" + // M[A] = 0 top of the stack is 0 as equal
// 		"(End)\n" // label End
// }

func Gt(counter int) string {
	return "@SP\n" + // A=0
		"M=M-1\n" + //point sp at top
		"A=M\n" + //set mem to sp
		"D=M\n" + // SAVE THE TOP OF THE STACK
		"A=A-1\n" + //POP MOVE A TO THE NEXT ITEM IN STACK
		"D=M-D\n" + //X-Y FOR THE 2 STACK NUMS
		"@IF_TRUE" + strconv.Itoa(counter) + "\n" + //jump label
		"D;JGT\n" + // JUMP IF D >0
		"D=0\n" + //if not equal set d=0
		"@END" + strconv.Itoa(counter) + "\n" + //SET SUMP LOACTION
		"D;JMP\n" + //JUMP TO THE END
		"(IF_TRUE" + strconv.Itoa(counter) + ")\n" + // true label
		"D=-1\n" + //SET D=-1 THOUGHT IT WOULD BE 0 BUT THE TESTS WANT -1
		"(END" + strconv.Itoa(counter) + ")" + "\n" +
		"@SP\n" +
		"A=M\n" +
		"A=A-1\n" +
		"M=D\n"
}
func Lt(counter int) string {
	return "@SP\n" + // A=0
		"M=M-1\n" + //point sp at top
		"A=M\n" + //set mem to sp
		"D=M\n" + // SAVE THE TOP OF THE STACK
		"A=A-1\n" + //POP MOVE A TO THE NEXT ITEM IN STACK
		"D=M-D\n" + //X-Y FOR THE 2 STACK NUMS
		"@IF_TRUE" + strconv.Itoa(counter) + "\n" + //jump label
		"D;JLT\n" + // JUMP IF D = 0
		"D=0\n" + //if not equal set d=0
		"@END" + strconv.Itoa(counter) + "\n" + //SET SUMP LOACTION
		"D;JMP\n" + //JUMP TO THE END
		"(IF_TRUE" + strconv.Itoa(counter) + ")\n" + // true label
		"D=-1\n" + //SET D=-1 THOUGHT IT WOULD BE 0 BUT THE TESTS WANT -1
		"(END" + strconv.Itoa(counter) + ")" + "\n" +
		"@SP\n" +
		"A=M\n" +
		"A=A-1\n" +
		"M=D\n"
}

//realized that eq implements just equal this is after a push  x and push y happened
// func eq(loc_x string, loc_y string, loc_jump string)string {
// 	return "@" + loc_x + "\nD = M\n@" + loc_y + "D = D - M\n"
// }
