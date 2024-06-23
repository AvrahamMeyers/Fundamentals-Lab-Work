package compilationengine

import (
	"fmt"
	"os"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

// Helper function that writes to a file
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

func (X *comp) Constructor(fileName string, folderpath string, err error) {
	file, err := os.OpenFile(fileName+".xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	X.file = file
	X.err = err
	X.tokenizer.Constructor(fileName, folderpath)
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
			X.CompileClassVarDec()
		}
		for X.tokenizer.KeyWord() == "constructor" || X.tokenizer.KeyWord() == "function" || X.tokenizer.KeyWord() == "method" {
			X.CompileSubroutine()
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
	helpWrite(X.file, "<classVarDec>\n", X.err, X.tabAmount)
	X.tabAmount += 1
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
	X.tabAmount -= 1
	helpWrite(X.file, "</classVarDec>\n", X.err, X.tabAmount)
}

// Compiles a complete method,
// function, or constructor.
func (X *comp) CompileSubroutine() {
	helpWrite(X.file, "<subroutineDec>\n", X.err, X.tabAmount)
	X.tabAmount += 1

	// ('constructor'|'function'|'method') ('void'|type)subroutineName'('parameterList')' subroutineBody
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//assumes next token is correct keyword void or type <keyword>
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//assumes next token is correct subroutine name <identifier>
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//assumes next toekn is correct <symbol>'('
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//call parameter list
	X.CompileParameterList()
	//assumes next token is correct <symbol ')'
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//subroutine body
	helpWrite(X.file, "<subroutineBody>\n", X.err, X.tabAmount)
	X.tabAmount += 1
	//assumes next token is correct <symbol> '{'
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//var declerations
	X.CompileVarDec()
	//statements
	X.CompileStatements()

	//assumes next token is correct <symbol> '}'
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//end subroutine body </subroutineBody>
	X.tabAmount -= 1
	helpWrite(X.file, "</subroutineBody>\n", X.err, X.tabAmount)

	X.tabAmount -= 1
	helpWrite(X.file, "</subroutineDec>\n", X.err, X.tabAmount)
}

// Compiles a (possibly empty) parameter list, not including the enclosing ‘‘()’’.
func (X *comp) CompileParameterList() {
	//not adding tabamount in this function as a test to see if we actually need tab amount
	//start param list
	helpWrite(X.file, "<parameterList>\n", X.err, X.tabAmount)

	//assumes the token expected is ) unless its a parameter
	for X.tokenizer.TokenType() == "symbol" && X.tokenizer.Symbol() == ")" {
		//type
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		//varname
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()

		//if there are more parameters if there aren't the for loop will be checked again and will break
		if X.tokenizer.Symbol() == "," {
			helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
			X.tokenizer.Advance()
		}
	}

	helpWrite(X.file, "</parameterList>\n", X.err, X.tabAmount)
}

// Compiles a var declaration.
func (X *comp) CompileVarDec() {
	for X.tokenizer.KeyWord() == "var" {
		helpWrite(X.file, "<varDec\n>", X.err, X.tabAmount)
		//var
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		//type
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		//varName
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		for X.tokenizer.Symbol() == "," {
			//,
			helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
			X.tokenizer.Advance()
			//varName
			helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
			X.tokenizer.Advance()
		}
		//;
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		helpWrite(X.file, "</varDec>\n", X.err, X.tabAmount)

	}
}

// Compiles a sequence of statements, not including the enclosing ‘‘{}’’.
func (X *comp) CompileStatements() {
	helpWrite(X.file, "<statements>\n", X.err, X.tabAmount)
	for X.tokenizer.KeyWord() == "let" ||
		X.tokenizer.KeyWord() == "if" ||
		X.tokenizer.KeyWord() == "while" ||
		X.tokenizer.KeyWord() == "do" ||
		X.tokenizer.KeyWord() == "return" {
		switch X.tokenizer.KeyWord() {
		case "let":
			X.CompileLet()
		case "if":
			X.CompileIf()
		case "while":
			X.CompileWhile()
		case "do":
			X.CompileDo()
		case "return":
			X.CompileReturn()
		default:
			fmt.Println("something unexpected happened in the Compile statements function")
		}
	}
	helpWrite(X.file, "</statements>\n", X.err, X.tabAmount)
}

// Compiles a do statement.
func (X *comp) CompileDo() {
	helpWrite(X.file, "<doStatement>\n", X.err, X.tabAmount)
	//do
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	X.CompileSubroutineCall()
	//symbol ;
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	helpWrite(X.file, "</doStatement>\n", X.err, X.tabAmount)
}

// Compiles a let statement.
func (X *comp) CompileLet() {
	helpWrite(X.file, "<letStatement>\n", X.err, X.tabAmount)
	//let
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	//varName
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	// [expression]?
	if X.tokenizer.Symbol() == "[" {
		//[
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		X.CompileExpression()
		//]
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		//=
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
		X.CompileExpression()
		//;
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
		X.tokenizer.Advance()
	}
	//symbol ;
	helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount)
	X.tokenizer.Advance()
	helpWrite(X.file, "</letStatement>\n", X.err, X.tabAmount)
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

// Compiles a subroutine call
func (X *comp) CompileSubroutineCall() {

}

// Compiles a term. This routine is faced with a slight difficulty
// when trying to decide between some of the alternative parsing
// rules. Specifically, if the current token is an identifier, the
// routine must distinguish between a variable, an array entry, and a
// subroutine call. A single look ahead token, which may be one
// of ‘‘[’’, ‘‘(’’, or ‘‘.’’ suffices to distinguish between the three
// possibilities. Any other token is not part of this term and
// should not be advanced over. Follows grammar:
// integerConstant | stringConstant | keywordConstant | varName | varName '[' expression ']' |
//
//	subroutineCall | '(' expression ')' | unaryOp term
func (X *comp) CompileTerm() {
	helpWrite(X.file, "<term>\n", X.err, X.tabAmount)
	X.tabAmount += 1

	// write the first (and maybe only) token including possibly varName, '(', or unaryOp
	helpWrite(X.file, X.tokenizer.FormatTokenString(), X.err, X.tabAmount)

	var firstType string = X.tokenizer.TokenType()
	var firstToken string = X.tokenizer.Token
	X.tokenizer.Advance()

	if firstType == "INT_CONST" || firstType == "STRING_CONST" || firstType == "KEYWORD" {
		// Nothing else needs to be done

	} else if firstType == "IDENTIFIER" { // varName | varName '[' expression ']' | subroutineCall

		if X.tokenizer.Token == "[" { // varName '[' expression ']'
			helpWrite(X.file, X.tokenizer.FormatTokenString(), X.err, X.tabAmount) // write the '['
			X.tokenizer.Advance()
			X.CompileExpression()
			helpWrite(X.file, X.tokenizer.FormatTokenString(), X.err, X.tabAmount) // write the ']'
			X.tokenizer.Advance()

		} else if X.tokenizer.Token == "(" || X.tokenizer.Token == "." { // subroutineCall
			X.CompileSubroutineCall()

		} // else { // just varName
		// Nothing else needs to be done, as the varName has already been written

	} else if firstType == "SYMBOL" { // '(' expression ')' | unaryOp term
		if firstToken == "(" { // '(' expression ')'
			X.CompileExpression()
			helpWrite(X.file, X.tokenizer.FormatTokenString(), X.err, X.tabAmount) // write the ')'
			X.tokenizer.Advance()

		} else { // unaryOp term
			X.CompileTerm()
		}
	}

	X.tabAmount -= 1
	helpWrite(X.file, "</term>\n", X.err, X.tabAmount)
}

// Compiles a (possibly empty) comma-separated list of expressions.
// Follows grammar: (expression (',' expression)* )?
func (X *comp) CompileExpressionList() {
	// As expressionLists always have a ')' after, check if it does, if so the expression list is empty
	if X.tokenizer.Symbol() == ")" {
		helpWrite(X.file, "<expressionList> </expressionList>\n", X.err, X.tabAmount)
		X.tokenizer.Advance()
		return
	}
	helpWrite(X.file, "<expressionList>\n", X.err, X.tabAmount)
	X.tabAmount += 1
	// expression
	X.CompileExpression()
	X.tokenizer.Advance()

	//(',' expression)*
	for X.tokenizer.Symbol() == "," {
		helpWrite(X.file, X.tokenizer.Token, X.err, X.tabAmount) // write the comma
		X.tokenizer.Advance()
		X.CompileExpression()
	}
	X.tabAmount -= 1
	helpWrite(X.file, "</expressionList>\n", X.err, X.tabAmount)
}
