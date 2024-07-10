package VMWriter

import (
	"fmt"
	"os"
)

type VMWriter struct {
	file *os.File
}

func (X *VMWriter) Constructor(fileName string, folderpath string) {
	file, err := os.OpenFile(fileName+".vm", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file in VM writer constructor")
		return
	}
	X.file = file
}

/*
Writes a VM push command.

Segment (CONST,
ARG, LOCAL,
STATIC, THIS,
THAT, POINTER,
TEMP)
Index (int)
*/
func (X *VMWriter) WritePush(segment string, index int) {

}

/*
Writes a VM pop command.

Segment (CONST,
ARG, LOCAL,
STATIC, THIS,
THAT, POINTER,
TEMP)
Index (int)
*/
func (X *VMWriter) WritePop(segment string, index int) {

}

/*
Writes a VM arithmetic
command.

command (ADD,
SUB, NEG, EQ, GT,
LT, AND, OR, NOT)
*/
func (X *VMWriter) WriteArithmetic(command string) {

}

/*
Writes a VM label command.
*/
func (X *VMWriter) WriteLabel(label string) {

}

/*
Writes a VM goto command.
*/
func (X *VMWriter) WriteGoto(label string) {

}

/*
Writes a VM If-goto command.
*/
func (X *VMWriter) WriteIf(label string) {

}

/*
Writes a VM call command

name (String)
nArgs (int)
*/
func (X *VMWriter) WriteCall(name string, nArgs int) {

}

/*
Writes a VM function command.

name (String)
nLocals (int)
*/
func (X *VMWriter) WriteFunction(name string, nLocals int) {

}

/*
Writes a VM return command.
*/
func (X *VMWriter) WriteReturn() {

}

/*
Closes the output file.
*/
func (X *VMWriter) Close() {
	X.file.Close()
}
