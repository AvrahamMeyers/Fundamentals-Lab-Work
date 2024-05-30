package CodeWriter

import "strconv"

func Add() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=D+M\n" // M[A] = D + M[A] (the value at the top of the stack)
}

func Sub() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=M-D\n" // M[A] = M[A] - D (the value at the top of the stack)
}

func Neg() string {
	return "@SP\n" +
		"A=M-1\n" +
		"M=-M\n"
}

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

func And() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=D&M\n" // M[A] = D and M[A] (the value at the top of the stack)

}

func Or() string {
	return "@SP\n" + // A = 0, the location of SP
		"M=M-1\n" + // SP-- (Move stack pointer one down)
		"A=M\n" + // A = SP (Save this pointer in A)
		"D=M\n" + // D = M[A] (the value at the top of the stack)
		"@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = M[A] - 1 (the location of the second of the stack)
		"M=D|M\n" // M[A] = D and M[A] (the value at the top of the stack)

}

func Not() string {
	return "@SP\n" + // A = 0, the location of SP
		"A=M-1\n" + // A = SP - 1 (address of the value at the top of the stack)
		"M=!M\n" // M[A] = not M[A] (the value at the top of the stack)
}
