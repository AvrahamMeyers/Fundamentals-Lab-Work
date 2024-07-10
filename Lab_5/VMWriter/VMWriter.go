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
	if segment == "CONST" {
		X.file.WriteString("push constant " + string(index) + "\n")
	} else if segment == "ARG" {
		X.file.WriteString("push argument " + string(index) + "\n")
	} else if segment == "LOCAL" {
		X.file.WriteString("push local " + string(index) + "\n")
	} else if segment == "STATIC" {
		X.file.WriteString("push static " + string(index) + "\n")
	} else if segment == "THIS" {
		X.file.WriteString("push this " + string(index) + "\n")
	} else if segment == "THAT" {
		X.file.WriteString("push that " + string(index) + "\n")
	} else if segment == "POINTER" {
		X.file.WriteString("push pointer " + string(index) + "\n")
	} else if segment == "TEMP" {
		X.file.WriteString("push temp " + string(index) + "\n")
	}
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
	if segment == "CONST" {
		X.file.WriteString("pop constant " + string(index) + "\n")
	} else if segment == "ARG" {
		X.file.WriteString("pop argument " + string(index) + "\n")
	} else if segment == "LOCAL" {
		X.file.WriteString("pop local " + string(index) + "\n")
	} else if segment == "STATIC" {
		X.file.WriteString("pop static " + string(index) + "\n")
	} else if segment == "THIS" {
		X.file.WriteString("pop this " + string(index) + "\n")
	} else if segment == "THAT" {
		X.file.WriteString("pop that " + string(index) + "\n")
	} else if segment == "POINTER" {
		X.file.WriteString("pop pointer " + string(index) + "\n")
	} else if segment == "TEMP" {
		X.file.WriteString("pop temp " + string(index) + "\n")
	}
}

/*
Writes a VM arithmetic
command.

command (ADD,
SUB, NEG, EQ, GT,
LT, AND, OR, NOT)
*/
func (X *VMWriter) WriteArithmetic(command string) {
	if command == "ADD" {
		X.file.WriteString("add\n")
	} else if command == "SUB" {
		X.file.WriteString("sub\n")
	} else if command == "NEG" {
		X.file.WriteString("neg\n")
	} else if command == "EQ" {
		X.file.WriteString("eq\n")
	} else if command == "GT" {
		X.file.WriteString("gt\n")
	} else if command == "LT" {
		X.file.WriteString("lt\n")
	} else if command == "AND" {
		X.file.WriteString("and\n")
	} else if command == "OR" {
		X.file.WriteString("or\n")
	} else if command == "NOT" {
		X.file.WriteString("not\n")
	}
}

/*
Writes a VM label command.
*/
func (X *VMWriter) WriteLabel(label string) {
	X.file.WriteString("label " + label + "\n")
}

/*
Writes a VM goto command.
*/
func (X *VMWriter) WriteGoto(label string) {
	X.file.WriteString("goto " + label + "\n")
}

/*
Writes a VM If-goto command.
*/
func (X *VMWriter) WriteIf(label string) {
	X.file.WriteString("if-goto " + label + "\n")
}

/*
Writes a VM call command

name (String)
nArgs (int)
*/
func (X *VMWriter) WriteCall(name string, nArgs int) {
	X.file.WriteString("call " + name + " " + string(nArgs) + "\n")
}

/*
Writes a VM function command.

name (String)
nLocals (int)
*/
func (X *VMWriter) WriteFunction(name string, nLocals int) {
	X.file.WriteString("function " + name + " " + string(nLocals) + "\n")
}

/*
Writes a VM return command.
*/
func (X *VMWriter) WriteReturn() {
	X.file.WriteString("return\n")
}

/*
Closes the output file.
*/
func (X *VMWriter) Close() {
	X.file.Close()
}
