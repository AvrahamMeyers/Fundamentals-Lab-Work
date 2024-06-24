package CompilationEngine

import (
	"fmt"
	"os"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

// Helper function that writes to a file
// remember the file needs to be open for append
func helpWrite(file *os.File, text string) {

	fmt.Println(text)
	_, err := (*file).WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// Holds information about the compiler
// Tokenizer object holds the current token
// file is the file that the xml will be written to
// tabAmount is the indentation level of the current line
type CompilationEngine struct {
	tokenizer Tokenizer.Tokenizer
	file      *os.File
}

func (X *CompilationEngine) Constructor(fileName string, folderpath string) {
	X.tokenizer.Constructor(fileName+".jack", folderpath)
	// outputFile, err := os.Create(fileName + "T.xml")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// final_str := X.tokenizer.TokenizeFile()
	// _, err = outputFile.WriteString(final_str)
	// if err != nil {
	// 	fmt.Println("Error appending to file:", err)
	// 	return
	// }
	file, err := os.OpenFile(fileName+"New"+".xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	X.file = file

	X.tokenizer.Advance()
	fmt.Println(X.tokenizer.HasMoreTokens())
	//A jack program will always begin with the word class
	X.CompileClass()

}

// Compiles a complete class.
func (X *CompilationEngine) CompileClass() {
	// class: 'class'className'{'classVarDec*subroutineDec*'}'
	if X.tokenizer.KeyWord() == "class" {
		helpWrite(X.file, "<class>\n")

		//'class'
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

		//className (identifier)
		if X.tokenizer.TokenType() != "IDENTIFIER" {
			//throw an error
			return
		}
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		//X.CompileTerm()
		//X.tokenizer.Advance()
		//symbol {
		if X.tokenizer.Symbol() != "{" {
			//throw an error
			return
		}
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		for X.tokenizer.KeyWord() == "static" || X.tokenizer.KeyWord() == "field" {
			X.CompileClassVarDec()
		}
		for X.tokenizer.KeyWord() == "constructor" || X.tokenizer.KeyWord() == "function" || X.tokenizer.KeyWord() == "method" {
			X.CompileSubroutine()
		}

		//at this point we have reached the end of the class and should only have '}' left
		// if X.tokenizer.Symbol() != "}" {
		// 	//throw an error
		// 	return
		// }
		helpWrite(X.file, X.tokenizer.FormatTokenString())

		helpWrite(X.file, "</class>\n")

	} else {
		fmt.Println("There's a problem with the file it does not begin with class")
	}
}

/*
Compiles a static declaration or
a field declaration.
*/
func (X *CompilationEngine) CompileClassVarDec() {
	helpWrite(X.file, "<classVarDec>\n")

	//it was already determined that the current token is static or field
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	//type
	if X.tokenizer.KeyWord() != "int" || X.tokenizer.KeyWord() != "char" || X.tokenizer.KeyWord() != "boolean" || X.tokenizer.TokenType() != "identifier" {
		//throw an error
		return
	}
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	if X.tokenizer.TokenType() != "identifier" {
		//throw an error
		return
	}
	X.CompileTerm()
	X.tokenizer.Advance()
	for X.tokenizer.Symbol() == "," {
		helpWrite(X.file, X.tokenizer.FormatTokenString())
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
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	helpWrite(X.file, "</classVarDec>\n")
}

// Compiles a complete method,
// function, or constructor.
func (X *CompilationEngine) CompileSubroutine() {
	helpWrite(X.file, "<subroutineDec>\n")

	// ('constructor'|'function'|'method') ('void'|type)subroutineName'('parameterList')' subroutineBody
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//assumes next token is correct keyword void or type <keyword>
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//assumes next token is correct subroutine name <identifier>
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//assumes next toekn is correct <symbol>'('
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//call parameter list
	X.CompileParameterList()
	//assumes next token is correct <symbol ')'
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//subroutine body
	helpWrite(X.file, "<subroutineBody>\n")

	//assumes next token is correct <symbol> '{'
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//var declerations
	X.CompileVarDec()
	//statements
	X.CompileStatements()

	//assumes next token is correct <symbol> '}'
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//end subroutine body </subroutineBody>

	helpWrite(X.file, "</subroutineBody>\n")

	helpWrite(X.file, "</subroutineDec>\n")
}

// Compiles a (possibly empty) parameter list, not including the enclosing ‘‘()’’.
func (X *CompilationEngine) CompileParameterList() {
	//not adding tabamount in this function as a test to see if we actually need tab amount
	//start param list
	helpWrite(X.file, "<parameterList>\n")

	//assumes the token expected is ) unless its a parameter
	for X.tokenizer.TokenType() == "symbol" && X.tokenizer.Symbol() == ")" {
		//type
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		//varname
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

		//if there are more parameters if there aren't the for loop will be checked again and will break
		if X.tokenizer.Symbol() == "," {
			helpWrite(X.file, X.tokenizer.FormatTokenString())
			X.tokenizer.Advance()
		}
	}

	helpWrite(X.file, "</parameterList>\n")
}

// Compiles a var declaration.
func (X *CompilationEngine) CompileVarDec() {
	for X.tokenizer.KeyWord() == "var" {
		helpWrite(X.file, "<varDec>\n")
		//var
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		//type
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		//varName
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		for X.tokenizer.Symbol() == "," {
			//,
			helpWrite(X.file, X.tokenizer.FormatTokenString())
			X.tokenizer.Advance()
			//varName
			helpWrite(X.file, X.tokenizer.FormatTokenString())
			X.tokenizer.Advance()
		}
		//;
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		helpWrite(X.file, "</varDec>\n")

	}
}

// Compiles a sequence of statements, not including the enclosing ‘‘{}’’.
func (X *CompilationEngine) CompileStatements() {
	helpWrite(X.file, "<statements>\n")
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
	helpWrite(X.file, "</statements>\n")
}

// Compiles a do statement.
func (X *CompilationEngine) CompileDo() {
	helpWrite(X.file, "<doStatement>\n")
	//do
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileSubroutineCall()
	//symbol ;
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	helpWrite(X.file, "</doStatement>\n")
}

// Compiles a let statement.
func (X *CompilationEngine) CompileLet() {
	helpWrite(X.file, "<letStatement>\n")
	//let
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//varName
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	// [expression]?
	if X.tokenizer.Symbol() == "[" {
		//[
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		X.CompileExpression()
		//]
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

	}
	//=
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileExpression()

	//symbol ;
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	helpWrite(X.file, "</letStatement>\n")
}

// Compiles a while statement.
func (X *CompilationEngine) CompileWhile() {
	helpWrite(X.file, "<whileStatement>\n")
	//while
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//symbol (
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileExpression()
	//symbol )
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//symbol {
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileStatements()
	// }
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	helpWrite(X.file, "</whileStatement>\n")
}

// Compiles a return statement.
func (X *CompilationEngine) CompileReturn() {
	helpWrite(X.file, "<returnStatement>\n")
	//return
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	if X.tokenizer.Symbol() != ";" {
		X.CompileExpression()
	}
	//symbol ;
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	helpWrite(X.file, "</returnStatement>\n")
}

// Compiles an if statement, possibly with a trailing else clause.
// Grammar: ifStatement: 'if' '(' expression ')' '{' statements '}'
// ('else' '{' statements '}')?
func (X *CompilationEngine) CompileIf() {
	helpWrite(X.file, "<ifStatement>\n")
	//if
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//symbol '('
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileExpression()
	//symbol ')'
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//symbol '{'
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileStatements()
	//symbol '}'
	helpWrite(X.file, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	if X.tokenizer.KeyWord() == "else" {
		//else
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		//symbol '{'
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		X.CompileStatements()
		//symbol '}'
		helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
	}
	helpWrite(X.file, "</ifStatement>\n")
}

// Compiles an expression
// Grammar: expression: term (op term)*
func (X *CompilationEngine) CompileExpression() {
	helpWrite(X.file, "<expression>\n")

	X.CompileTerm()
	if X.tokenizer.TokenType() == "SYMBOL" {
		var token string = X.tokenizer.Symbol()
		for token == "+" || token == "-" || token == "*" || token == "/" || token == "&amp;" || token == "|" || token == "&lt;" || token == "&gt;" || token == "=" {
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: op
			X.tokenizer.Advance()
			X.CompileTerm()
			token = X.tokenizer.Symbol()
		}
	}

	helpWrite(X.file, "</expression>\n")
}

// Compiles a subroutine call, not a seperate function, doesn't have its own tags
// Grammar: subroutineCall: subroutineName '(' expressionList ')' | (className |
// varName) '.' subroutineName '(' expressionList ')'
func (X *CompilationEngine) CompileSubroutineCall() {
	helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: subroutineName or className or varName
	X.tokenizer.Advance()
	if X.tokenizer.Token == "(" {
		helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: '('
		X.tokenizer.Advance()
		X.CompileExpressionList()
	} else {
		helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: '.'
		X.tokenizer.Advance()
		helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: subroutineName
		X.tokenizer.Advance()
		helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: '('
		X.tokenizer.Advance()
		X.CompileExpressionList() // write: expressionList
	}
	helpWrite(X.file, X.tokenizer.FormatTokenString()) // write: ')'
	X.tokenizer.Advance()
}

// Compiles a term. This routine is faced with a slight difficulty
// when trying to decide between some of the alternative parsing
// rules. Specifically, if the current token is an identifier, the
// routine must distinguish between a variable, an array entry, and a
// subroutine call. A single look ahead token, which may be one
// of ‘‘[’’, ‘‘(’’, or ‘‘.’’ suffices to distinguish between the three
// possibilities. Any other token is not part of this term and
// should not be advanced over. Follows grammar:
// integerConstant | stringConstant | keywordConstant | varName | varName '[' expression ']'
// | subroutineCall | '(' expression ')' | unaryOp term
func (X *CompilationEngine) CompileTerm() {
	helpWrite(X.file, "<term>\n")

	var firstType string = X.tokenizer.TokenType()

	if firstType == "INT_CONST" || firstType == "STRING_CONST" || firstType == "KEYWORD" {
		helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the constant or keyword
		X.tokenizer.Advance()

	} else if firstType == "IDENTIFIER" { // varName | varName '[' expression ']' | subroutineCall
		var lookahead Tokenizer.Tokenizer = X.tokenizer
		lookahead.Advance()

		if lookahead.Token == "[" { // varName '[' expression ']'
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the varName
			X.tokenizer.Advance()
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the '['
			X.tokenizer.Advance()
			X.CompileExpression()
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the ']'
			X.tokenizer.Advance()

		} else if lookahead.Token == "(" || lookahead.Token == "." { // subroutineCall
			X.CompileSubroutineCall()

		} else { // just varName
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the varName
			X.tokenizer.Advance()
		}

	} else if firstType == "SYMBOL" { // '(' expression ')' | unaryOp term
		if X.tokenizer.Token == "(" { // '(' expression ')'
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the '('
			X.tokenizer.Advance()
			X.CompileExpression()
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the ')'
			X.tokenizer.Advance()

		} else { // unaryOp term
			helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the unaryOp
			X.tokenizer.Advance()
			X.CompileTerm()
		}
	}

	helpWrite(X.file, "</term>\n")
}

// Compiles a (possibly empty) comma-separated list of expressions.
// Follows grammar: (expression (',' expression)* )?
func (X *CompilationEngine) CompileExpressionList() {
	// As expressionLists always have a ')' after, check if it does, if so the expression list is empty
	if X.tokenizer.Symbol() == ")" {
		helpWrite(X.file, "<expressionList> </expressionList>\n")
		X.tokenizer.Advance()
		return
	}
	helpWrite(X.file, "<expressionList>\n")

	// expression
	X.CompileExpression()

	//(',' expression)*
	for X.tokenizer.Symbol() == "," {
		helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the comma
		X.tokenizer.Advance()
		X.CompileExpression()
	}

	helpWrite(X.file, "</expressionList>\n")
}
