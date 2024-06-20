package compilationengine

import (
	"fmt"
	"os"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

// Function that writes to a file
// remember the file needs to be open for append
func helpWrite(file *os.File, text string, err error, tab int) {
	tabs := ""
	for i := 0; i < tab; i++ {
		tabs += "\t"
	}
	_, err = file.WriteString(tabs + text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// Holds information about the compiler
// Tokenizer object holds the current token
// file is the file that the xml will be written to
// tabAmount is the indentation level of the current line
type comp struct {
	tokenizer Tokenizer.Tokenizer
	file      *os.File
	err       error
	tabAmount int
}

func (X *comp) Constructor(fileName string, err error) {
	file, err := os.OpenFile(fileName+".xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	X.file = file
	X.err = err
	X.tokenizer.Constructor(fileName)
	//A jack program will always begin with the word class
	X.CompileClass()
	X.tabAmount = 0
}

// Compiles a complete class.
func (X *comp) CompileClass() {
	// class: 'class'className'{'classVarDec*subroutineDec*'}'
	if X.tokenizer.TokenType() == "class" {
		helpWrite(X.file, "<class>\n", X.err, X.tabAmount)
		X.tabAmount += 1
		//'class'
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()

		//className (identifier)
		if X.tokenizer.TokenType() != "identifier" {
			//throw an error
			return
		}
		X.CompileTerm()
		X.tokenizer.Advance()
		//symbol {
		if X.tokenizer.Symbol() != "{" {
			//throw an error
			return
		}
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		for X.tokenizer.KeyWord() == "static" || X.tokenizer.KeyWord() == "field" {
			helpWrite(X.file, "<classVarDec>\n", X.err, X.tabAmount)
			X.tabAmount += 1
			X.CompileClassVarDec()
			X.tabAmount -= 1
			helpWrite(X.file, "</classVarDec>\n", X.err, X.tabAmount)
		}
		for X.tokenizer.KeyWord() == "constructor" || X.tokenizer.KeyWord() == "function" || X.tokenizer.KeyWord() == "method" {
			helpWrite(X.file, "<subroutineDec>\n", X.err, X.tabAmount)
			X.tabAmount += 1
			X.CompileSubroutine()
			X.tabAmount -= 1
			helpWrite(X.file, "</subroutineDec>\n", X.err, X.tabAmount)
		}

		//at this point we have reached the end of the class and should only have '}' left
		if X.tokenizer.Symbol() != "}" {
			//throw an error
			return
		}
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tabAmount -= 1
		helpWrite(X.file, "</class>\n", X.err, X.tabAmount)

	} else {
		fmt.Println("There's a problem with the file it does not begin with class")
	}
}

/*
Compiles a static declaration or
a field declaration.
*/
func (X *comp) CompileClassVarDec() {
	//it was already determined that the current token is static or field
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()

	//type
	if X.tokenizer.KeyWord() != "int" || X.tokenizer.KeyWord() != "char" || X.tokenizer.KeyWord() != "boolean" || X.tokenizer.TokenType() != "identifier" {
		//throw an error
		return
	}
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	if X.tokenizer.TokenType() != "identifier" {
		//throw an error
		return
	}
	X.CompileTerm()
	X.tokenizer.Advance()
	for X.tokenizer.Symbol() == "," {
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		if X.tokenizer.TokenType() != "identifier" {
			//throw an error
			return
		}
		X.CompileTerm()
		X.tokenizer.Advance()
	}
	if X.tokenizer.Symbol() != "," {
		//throw an error
		return
	}
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()

}

// Compiles a complete method,
// function, or constructor.
func (X *comp) CompileSubroutine() {

}

// Compiles a (possibly empty) parameter list, not including the enclosing ‘‘()’’.
func (X *comp) CompileParameterList() {

}

// Compiles a var declaration.
func (X *comp) CompileVarDec() {

}

// Compiles a sequence of statements, not including the enclosing ‘‘{}’’.
func (X *comp) CompileStatements() {

}

// Compiles a do statement.
func (X *comp) CompileDo() {

}

// Compiles a let statement.
func (X *comp) CompileLet() {

}

// Compiles a while statement.
func (X *comp) CompileWhile() {

}

// Compiles a return statement.
func (X *comp) CompileReturn() {

}

// Compiles an if statement, possibly with a trailing else clause.
func (X *comp) CompileIf() {

}

// Compiles an expression
func (X *comp) CompileExpression() {

}

// Compiles a term. This routine is faced with a slight difficulty
// when trying to decide between some of the alternative parsing
// rules. Specifically, if the current token is an identifier, the routine
// routine must distinguish between a variable, an array entry, and a
// subroutine call. A single look ahead token, which may be one
// of ‘‘[’’, ‘‘(’’, or ‘‘.’’ suffices to distinguish between the three
// possibilities. Any other token is not part of this term and
// should not be advanced over.*/
func (X *comp) CompileTerm() {

}

// Compiles a (possibly empty) comma-separated list of expressions.
func (X *comp) CompileExpressionList() {
}
