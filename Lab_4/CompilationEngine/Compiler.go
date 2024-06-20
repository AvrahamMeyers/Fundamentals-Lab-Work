package compilationengine

import (
	"fmt"
	"os"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

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

type comp struct {
	token     Tokenizer.Tokenizer
	write     *os.File
	err       error
	tabAmount int
}

func (X *comp) Constructor(fileName string, err error) {
	file, err := os.OpenFile(fileName+".xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	X.write = file
	X.err = err
	X.token.Constructor(fileName)
	//A jack program will always begin with the word class
	X.CompileClass()
	X.tabAmount = 0
}

// Compiles a complete class.
func (X *comp) CompileClass() {
	// class: 'class'className'{'classVarDec*subroutineDec*'}'
	if X.token.TokenType() == "class" {
		helpWrite(X.write, "<class>\n", X.err, X.tabAmount)
		X.tabAmount += 1
		//'class'
		helpWrite(X.write, X.token.Token, X.err, X.tabAmount)
		X.token.Advance()

		//className (identifier)
		if X.token.TokenType() != "identifier" {
			//throw an error
			return
		}
		X.CompileTerm()
		X.token.Advance()
		//symbol {
		if X.token.Symbol() != "{" {
			//throw an error
			return
		}
		helpWrite(X.write, X.token.Token, X.err, X.tabAmount)
		X.token.Advance()
		for X.token.KeyWord() == "static" || X.token.KeyWord() == "field" {
			helpWrite(X.write, "<classVarDec>\n", X.err, X.tabAmount)
			X.tabAmount += 1
			X.CompileClassVarDec()
			X.tabAmount -= 1
			helpWrite(X.write, "</classVarDec>\n", X.err, X.tabAmount)
		}
		for X.token.KeyWord() == "constructor" || X.token.KeyWord() == "function" || X.token.KeyWord() == "method" {
			helpWrite(X.write, "<subroutineDec>\n", X.err, X.tabAmount)
			X.tabAmount += 1
			X.CompileSubroutine()
			X.tabAmount -= 1
			helpWrite(X.write, "</subroutineDec>\n", X.err, X.tabAmount)
		}

		//at this point we have reached the end of the class and should only have '}' left
		if X.token.Symbol() != "}" {
			//throw an error
			return
		}
		helpWrite(X.write, X.token.Token, X.err, X.tabAmount)
		X.tabAmount -= 1
		helpWrite(X.write, "</class>\n", X.err, X.tabAmount)

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
	helpWrite(X.write, X.token.Token, X.err, X.tabAmount)
	X.token.Advance()
	if X.token.TokenType() != "identifier" {
		//throw an error
		return
	}
	X.CompileTerm()
	X.token.Advance()
	for X.token.Symbol() == "," {
		helpWrite(X.write, X.token.Token, X.err, X.tabAmount)
		X.token.Advance()
		if X.token.TokenType() != "identifier" {
			//throw an error
			return
		}
		X.CompileTerm()
		X.token.Advance()
	}
	if X.token.Symbol() != "," {
		//throw an error
		return
	}
	helpWrite(X.write, X.token.Token, X.err, X.tabAmount)
	X.token.Advance()

}

func (X *comp) CompileSubroutine() {
	//Compiles a complete method,
	//function, or constructor.
}

func (X *comp) CompileParameterList() {
	/*Compiles a (possibly empty)
	parameter list, not including the
	enclosing ‘‘()’’..*/
}

func (X *comp) CompileVarDec() {
	/* Compiles a var declaration..*/
}

func (X *comp) CompileStatements() {
	/* Compiles a sequence of state
	ments, not including the
	 enclosing ‘‘{}’’..*/
}

func (X *comp) CompileDo() {
	/* Compiles a do statement..*/
}

func (X *comp) CompileLet() {
	/*  Compiles a let statement..*/
}
func (X *comp) CompileWhile() {
	/* Compiles a while statement..*/
}
func (X *comp) CompileReturn() {
	/*  Compiles a return statement.*/
}
func (X *comp) CompileIf() {
	/*  Compiles an if statement, pos
	sibly with a trailing else clause.T.*/
}
func (X *comp) CompileExpression() {
	/* Compiles an expression..*/
}
func (X *comp) CompileTerm() {
	/*Compiles a term. This routine is
	 faced with a slight difficulty
	 when trying to decide between
	 some of the alternative parsing
	 rules. Specifically, if the current
	 token is an identifier, the routine
	 must distinguish between a
	 variable, an array entry, and a
	 subroutine call. A single look
	ahead token, which may be one
	 of ‘‘[’’, ‘‘(’’, or ‘‘.’’ suffices to dis
	tinguish between the three possi
	bilities. Any other token is not
	 part of this term and should not
	 be advanced over.*/
}
func (X *comp) CompileExpressionList() {
	/*Compiles a (possibly empty)
	comma-separated list of
	expressions.*/
}
