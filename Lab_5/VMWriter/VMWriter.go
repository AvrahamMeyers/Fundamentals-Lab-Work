package VMWriter

import (
	"fmt"
	"os"
)

type VMWriter struct {
	file *os.File
}

func (X *VMWriter) Constructor(fileName string, folderpath string) {
	filePath := folderpath + "/" + fileName + ".vm"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	segment = convert_kind_to_segment(segment)

	segment = convert_segment_to_vm(segment)

	if segment != "ERROR" {
		X.file.WriteString("push " + segment + " " + fmt.Sprint(index) + "\n")
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
	segment = convert_kind_to_segment(segment)

	segment = convert_segment_to_vm(segment)

	if segment != "ERROR" {
		X.file.WriteString("pop " + segment + " " + fmt.Sprint(index) + "\n")
	}
}

func convert_segment_to_vm(segment string) string {
	switch segment {
	case "CONST":
		return "constant"
	case "ARG":
		return "argument"
	case "LOCAL":
		return "local"
	case "STATIC":
		return "static"
	case "THIS":
		return "this"
	case "THAT":
		return "that"
	case "POINTER":
		return "pointer"
	case "TEMP":
		return "temp"
	default:
		return "ERROR"
	}
}

/*
Writes a VM arithmetic
command.

command (ADD, MUL,
SUB, NEG, EQ, GT,
LT, AND, OR, NOT)
*/
func (X *VMWriter) WriteArithmetic(command string) {
	switch command {
	case "ADD":
		X.file.WriteString("add\n")
	case "SUB":
		X.file.WriteString("sub\n")
	case "NEG":
		X.file.WriteString("neg\n")
	case "EQ":
		X.file.WriteString("eq\n")
	case "GT":
		X.file.WriteString("gt\n")
	case "LT":
		X.file.WriteString("lt\n")
	case "AND":
		X.file.WriteString("and\n")
	case "OR":
		X.file.WriteString("or\n")
	case "NOT":
		X.file.WriteString("not\n")
	case "MUL":
		X.file.WriteString("call Math.multiply 2\n")
	case "DIV":
		X.file.WriteString("call Math.divide 2\n")
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
	X.file.WriteString("call " + name + " " + fmt.Sprint(nArgs) + "\n")
}

/*
Writes a VM function command.

name (String)
nLocals (int)
*/
func (X *VMWriter) WriteFunction(name string, nLocals int) {
	X.file.WriteString("function " + name + " " + fmt.Sprint(nLocals) + "\n")
}

/*
Writes a VM return command.
*/
func (X *VMWriter) WriteReturn() {
	X.file.WriteString("return\n")
}

func (X *VMWriter) WriteMemAlloc(amount int) {
	X.WritePush("CONST", amount)
	X.file.WriteString("call Memory.alloc 1\n")
}

/*
Closes the output file.
*/
func (X *VMWriter) Close() {
	X.file.Close()
}

func convert_kind_to_segment(kind string) string {
	switch kind {
	case "CONST", "ARG", "LOCAL", "STATIC", "THIS", "THAT", "POINTER", "TEMP":
		return kind
	case "field":
		return "THIS"
	case "VAR":
		return "LOCAL"
	case "static":
		return "STATIC"
	default:
		return "ERROR"
	}
}
