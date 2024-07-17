package CompilationEngine

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_5/SymbolTable"
	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_5/Tokenizer"
	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_5/VMWriter"
)

// Helper function that writes to a file
// remember the file needs to be open for append
func helpWrite(file *os.File, text string) {

	//fmt.Println(text)
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
	tokenizer    Tokenizer.Tokenizer
	symbolTable  SymbolTable.SymbolTable
	xmlFile      *os.File
	vmwriter     VMWriter.VMWriter
	filename     string
	labelCounter int
}

func (X *CompilationEngine) Constructor(fileName string, folderpath string) {
	X.tokenizer.Constructor(fileName+".jack", folderpath)

	X.vmwriter.Constructor(fileName, folderpath)
	X.filename = fileName

	X.labelCounter = 0
	// outputFile, err := os.Create(folderpath + "/" + fileName + "T.xml")
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
	file, err := os.OpenFile(folderpath+"/"+fileName+".xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file in compiler constructor")
		return
	}
	defer file.Close()
	X.xmlFile = file

	X.tokenizer.Advance()
	//fmt.Println(X.tokenizer.HasMoreTokens())
	//A jack program will always begin with the word class
	X.CompileClass()

}

// Compiles a complete class.
// class grammar: 'class'className'{'classVarDec*subroutineDec*'}'
func (X *CompilationEngine) CompileClass() {
	if X.tokenizer.KeyWord() == "class" {
		// Make symbol tables for this class
		X.symbolTable.Constructor()
		helpWrite(X.xmlFile, "<class>\n")

		//'class'
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

		//className (identifier)
		if X.tokenizer.TokenType() != "IDENTIFIER" {
			//throw an error
			return
		}
		//class name
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

		//symbol {
		if X.tokenizer.Symbol() != "{" {
			//throw an error
			return
		}
		//{
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

		//variable decl.
		for X.tokenizer.KeyWord() == "static" || X.tokenizer.KeyWord() == "field" {
			X.CompileClassVarDec()
		}
		for X.tokenizer.KeyWord() == "constructor" || X.tokenizer.KeyWord() == "function" || X.tokenizer.KeyWord() == "method" {
			X.CompileSubroutine()
		}

		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())

		helpWrite(X.xmlFile, "</class>\n")

	} else {
		fmt.Println("There's a problem with the file it does not begin with class")
	}
}

/*
Compiles a static declaration or
a field declaration.
Grammar:  ('static'|'field') type varName (',' varName)* ';'
*/
func (X *CompilationEngine) CompileClassVarDec() {
	helpWrite(X.xmlFile, "<classVarDec>\n")

	//it was already determined that the current token is static or field
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write static or field
	var itsKind string = X.tokenizer.Token

	X.tokenizer.Advance()

	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the type
	var itsType string = X.tokenizer.Token
	X.tokenizer.Advance()

	// helpWrite(X.file, X.tokenizer.FormatTokenString()) // write varname
	X.symbolTable.Define(X.tokenizer.Token, itsType, itsKind)
	helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, true))

	X.tokenizer.Advance()

	for X.tokenizer.Token == "," {
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the ","
		X.tokenizer.Advance()

		// helpWrite(X.file, X.tokenizer.FormatTokenString()) // write varname
		X.symbolTable.Define(X.tokenizer.Token, itsType, itsKind)
		helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, true))
		X.tokenizer.Advance()
	}

	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the ";"
	X.tokenizer.Advance()

	helpWrite(X.xmlFile, "</classVarDec>\n")
}

// Compiles a complete method,
// function, or constructor.
// Grammar: ('constructor'|'function'|'method') ('void'|type)subroutineName'('parameterList')' subroutineBody
func (X *CompilationEngine) CompileSubroutine() {

	subroutineType := X.tokenizer.KeyWord()

	helpWrite(X.xmlFile, "<subroutineDec>\n")

	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	//assumes next token is correct keyword void or type <keyword>
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	label := ""
	//assumes next token is correct subroutine name <identifier>
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())

	label = X.filename + "." + X.tokenizer.Token

	X.tokenizer.Advance()

	//assumes next token is correct <symbol>'('
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//subroutine body
	helpWrite(X.xmlFile, "<subroutineBody>\n")
	X.symbolTable.StartSubroutine()
	//call parameter list
	X.CompileParameterList()
	//declare the function label

	//assumes next token is correct symbol ')'
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	//assumes next token is correct <symbol> '{'
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	//var declerations
	X.CompileVarDec()
	X.vmwriter.WriteFunction(label, X.symbolTable.VarCount("VAR"))

	if subroutineType == "constructor" {
		filedsCount := X.symbolTable.VarCount("field")
		X.vmwriter.WriteMemAlloc(filedsCount)
		X.vmwriter.WritePop("POINTER", 0)
	} else if subroutineType == "method" {
		X.vmwriter.WritePush("ARG", 0)
		X.vmwriter.WritePop("POINTER", 0)
	}

	fmt.Println("VarDec done " + label)
	fmt.Println(X.symbolTable.VarCount("VAR"))
	//statements
	X.CompileStatements()
	fmt.Println("Statements done " + label)

	//assumes next token is correct <symbol> '}'
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//end subroutine body </subroutineBody>

	helpWrite(X.xmlFile, "</subroutineBody>\n")

	helpWrite(X.xmlFile, "</subroutineDec>\n")
}

// Compiles a (possibly empty) parameter list, not including the enclosing ‘‘()’’.
// Grammar: parameterList: ((type varName) (',' type varName)*)?
func (X *CompilationEngine) CompileParameterList() int {

	// As expressionLists always have a ')' after, check if it does, if so the expression list is empty
	if X.tokenizer.Symbol() == ")" {
		helpWrite(X.xmlFile, "<parameterList> </parameterList>\n")
		// no need to advacnce, the ') is part of the caller of parameterList
		return 0
	}

	helpWrite(X.xmlFile, "<parameterList>\n") //start param list

	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the type
	var itsType string = X.tokenizer.Token
	X.tokenizer.Advance()
	i := 1
	// helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the varName
	X.symbolTable.Define(X.tokenizer.Token, itsType, "ARG")
	helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, true))
	X.tokenizer.Advance()

	for X.tokenizer.Symbol() == "," {
		i += 1
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the comma
		X.tokenizer.Advance()

		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the type
		itsType = X.tokenizer.Token
		X.tokenizer.Advance()

		// helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the varName
		X.symbolTable.Define(X.tokenizer.Token, itsType, "ARG")
		helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, true))
		X.tokenizer.Advance()
	}

	helpWrite(X.xmlFile, "</parameterList>\n")
	return i
}

// Compiles a var declaration.
// Original grammar: 'var' type varName (',' varName)* ';'
// Grammar changed to: ('var' type varName (',' varName)* ';')*
func (X *CompilationEngine) CompileVarDec() {
	for X.tokenizer.KeyWord() == "var" {
		helpWrite(X.xmlFile, "<varDec>\n")
		//var
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

		//type
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		var itsType string = X.tokenizer.Token
		X.tokenizer.Advance()

		//varName
		// helpWrite(X.file, X.tokenizer.FormatTokenString())
		X.symbolTable.Define(X.tokenizer.Token, itsType, "VAR")
		helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, true))
		X.tokenizer.Advance()

		for X.tokenizer.Symbol() == "," {
			//,
			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
			X.tokenizer.Advance()

			//varName
			// helpWrite(X.file, X.tokenizer.FormatTokenString())
			X.symbolTable.Define(X.tokenizer.Token, itsType, "VAR")
			helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, true))
			X.tokenizer.Advance()
		}
		//;
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		helpWrite(X.xmlFile, "</varDec>\n")
	}
}

// Compiles a sequence of statements, not including the enclosing ‘‘{}’’.
func (X *CompilationEngine) CompileStatements() {

	if X.tokenizer.Token == "}" {
		helpWrite(X.xmlFile, "<statements> </statements>\n")
		return
	}

	helpWrite(X.xmlFile, "<statements>\n")
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
	helpWrite(X.xmlFile, "</statements>\n")
}

// Compiles a do statement.
// 'do' subroutineCall ';'
func (X *CompilationEngine) CompileDo() {
	helpWrite(X.xmlFile, "<doStatement>\n")
	//do
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	X.CompileSubroutineCall()
	//symbol ;
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	helpWrite(X.xmlFile, "</doStatement>\n")
	X.vmwriter.WritePop("TEMP", 0)
}

// Compiles a let statement.
func (X *CompilationEngine) CompileLet() {
	helpWrite(X.xmlFile, "<letStatement>\n")
	//let
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	//varName
	// helpWrite(X.file, X.tokenizer.FormatTokenString())
	kind := X.symbolTable.KindOf(X.tokenizer.Token)
	num := X.symbolTable.IndexOf(X.tokenizer.Token)

	helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, false))
	X.tokenizer.Advance()

	isarray := false
	// [expression]?
	if X.tokenizer.Symbol() == "[" { //TODO: handle array
		isarray = true
		//[
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		X.CompileExpression()

		X.vmwriter.WritePush(kind, num) // push the base address

		X.vmwriter.WriteArithmetic("ADD") // the index with the base address
		X.vmwriter.WritePop("POINTER", 1) //SAVE X[NUMBER] TO POINTER 1
		//]
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()

	}
	//=
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileExpression()

	if isarray {
		// Pop the result into the array element (that 0)
		X.vmwriter.WritePop("THAT", 0)
	} else {
		X.vmwriter.WritePop(kind, num) //save the value to the appropriate location
	}
	//symbol ;
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	helpWrite(X.xmlFile, "</letStatement>\n")
}

// Compiles a while statement.
func (X *CompilationEngine) CompileWhile() {
	ThisScope := X.labelCounter + 1
	X.labelCounter++
	helpWrite(X.xmlFile, "<whileStatement>\n")
	//while
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	begin_loop_label := "LOOP" + strconv.Itoa(ThisScope)
	X.vmwriter.WriteLabel(begin_loop_label)

	//symbol (
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileExpression()
	//symbol )
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())

	end_label_loop := "END" + strconv.Itoa(ThisScope)
	X.vmwriter.WriteArithmetic("NOT")
	X.vmwriter.WriteIf(end_label_loop)

	X.tokenizer.Advance()
	//symbol {
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileStatements()

	X.vmwriter.WriteGoto(begin_loop_label)
	// }
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()

	helpWrite(X.xmlFile, "</whileStatement>\n")

	X.vmwriter.WriteLabel(end_label_loop)
}

// Compiles a return statement.
func (X *CompilationEngine) CompileReturn() {
	helpWrite(X.xmlFile, "<returnStatement>\n")
	//return
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	if X.tokenizer.Symbol() != ";" {
		X.CompileExpression()
	} else {
		X.vmwriter.WritePush("CONST", 0)
	}
	//symbol ;
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.vmwriter.WriteReturn()
	helpWrite(X.xmlFile, "</returnStatement>\n")
}

// Compiles an if statement, possibly with a trailing else clause.
// Grammar: ifStatement: 'if' '(' expression ')' '{' statements '}'
// ('else' '{' statements '}')?
func (X *CompilationEngine) CompileIf() {
	ThisScope := X.labelCounter + 1
	X.labelCounter++
	helpWrite(X.xmlFile, "<ifStatement>\n")
	//if
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	//symbol '('
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileExpression()
	else_label := "ELSE" + strconv.Itoa(ThisScope)
	X.vmwriter.WriteArithmetic("NOT")
	X.vmwriter.WriteIf(else_label)
	//symbol ')'
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())

	X.tokenizer.Advance()
	//symbol '{'
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.tokenizer.Advance()
	X.CompileStatements()
	end_label := "END" + strconv.Itoa(ThisScope)
	X.vmwriter.WriteGoto(end_label)
	//symbol '}'
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
	X.vmwriter.WriteLabel(else_label)
	X.tokenizer.Advance()
	if X.tokenizer.KeyWord() == "else" {
		//else
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		//symbol '{'
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
		X.CompileStatements()
		//symbol '}'
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString())
		X.tokenizer.Advance()
	}
	helpWrite(X.xmlFile, "</ifStatement>\n")
	X.vmwriter.WriteLabel(end_label)

}

// Compiles an expression
// Grammar: expression: term (op term)*
func (X *CompilationEngine) CompileExpression() {
	helpWrite(X.xmlFile, "<expression>\n")

	X.CompileTerm()
	if X.tokenizer.TokenType() == "SYMBOL" {
		var token string = X.tokenizer.Symbol()
		for token == "+" || token == "-" || token == "*" || token == "/" || token == "&" || token == "|" || token == "<" || token == ">" || token == "=" {
			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: op
			X.tokenizer.Advance()
			X.CompileTerm()
			token_type := convert_binary_token(token)
			X.vmwriter.WriteArithmetic(token_type)
			token = X.tokenizer.Symbol()
		}
	}

	helpWrite(X.xmlFile, "</expression>\n")
}

// Compiles a subroutine call, not a seperate function, doesn't have its own tags
// Grammar: subroutineCall: subroutineName '(' expressionList ')' | (className |
// varName) '.' subroutineName '(' expressionList ')'
func (X *CompilationEngine) CompileSubroutineCall() {
	i := 0 //how many argumants
	callName := ""
	callName = X.tokenizer.Token
	if X.symbolTable.KindOf(X.tokenizer.Token) == "NONE" {
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: subroutineName or className
		callName = X.tokenizer.Token
	} else {
		//X.symbolTable.IndexOf(callName)
		helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, false)) // write: varName
		objectName := X.tokenizer.Token
		kind := X.symbolTable.KindOf(objectName)
		typeName := X.symbolTable.TypeOf(objectName)
		index := X.symbolTable.IndexOf(objectName)
		i += 1 // One argument for the object reference
		X.vmwriter.WritePush(kind, index)

		callName = typeName // The type name becomes part of the call name
	}
	//beforeDot := X.tokenizer.Token
	X.tokenizer.Advance()
	if X.tokenizer.Token == "(" {
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: '('
		X.tokenizer.Advance()
		callName = X.filename + "." + callName
		X.vmwriter.WritePush("POINTER", 0)
		i += X.CompileExpressionList() + 1

	} else {
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: '.'
		X.tokenizer.Advance()
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: subroutineName
		nameOfFunc := X.tokenizer.Token
		X.tokenizer.Advance()
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: '('
		X.tokenizer.Advance()
		i += X.CompileExpressionList() // write: expressionList
		callName = callName + "." + nameOfFunc
	}
	helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write: ')'
	X.tokenizer.Advance()
	// for j := i - 1; j >= 0; j-- {
	// 	X.vmwriter.WritePop("ARG", j)
	// }
	X.vmwriter.WriteCall(callName, i)
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
	helpWrite(X.xmlFile, "<term>\n")

	var firstType string = X.tokenizer.TokenType()

	if firstType == "INT_CONST" || firstType == "STRING_CONST" || firstType == "KEYWORD" {
		if firstType == "INT_CONST" {
			i, _ := strconv.Atoi(X.tokenizer.Token)
			X.vmwriter.WritePush("CONST", i)
		} else if firstType == "KEYWORD" {
			switch X.tokenizer.KeyWord() {
			case "true":
				X.vmwriter.WritePush("CONST", 1)
				X.vmwriter.WriteArithmetic("NEG")
			case "false", "null":
				X.vmwriter.WritePush("CONST", 0)
			case "this":
				X.vmwriter.WritePush("POINTER", 0)
			}
		} else if firstType == "STRING_CONST" {
			str := X.tokenizer.Token
			X.vmwriter.WritePush("CONST", len(str))
			X.vmwriter.WriteCall("String.new", 1)
			for i := 0; i < len(str); i++ {
				X.vmwriter.WritePush("CONST", int(str[i]))
				X.vmwriter.WriteCall("String.appendChar", 2)
			}
		}
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the constant or keyword
		X.tokenizer.Advance()

	} else if firstType == "IDENTIFIER" { // varName | varName '[' expression ']' | subroutineCall
		var lookahead Tokenizer.Tokenizer = X.tokenizer
		lookahead.Advance()

		if lookahead.Token == "[" { // varName '[' expression ']'
			// helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the varName
			helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, false))

			varname := X.tokenizer.Token //address 0 of the array we need to add to that
			addzero := X.symbolTable.IndexOf(varname)
			argument := X.symbolTable.TypeOf(varname) //the rest of the array has to be addressed
			//push the address of the 0 pointer
			X.tokenizer.Advance()

			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the '['
			X.tokenizer.Advance()

			X.CompileExpression()

			X.vmwriter.WritePush(argument, addzero)

			X.vmwriter.WriteArithmetic("ADD")
			X.vmwriter.WritePop("POINTER", 1)
			X.vmwriter.WritePush("THAT", 0)
			/////////////////X.vmwriter.WritePop("TEMP", 0)
			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the ']'
			X.tokenizer.Advance()

		} else if lookahead.Token == "(" || lookahead.Token == "." { // subroutineCall
			X.CompileSubroutineCall()

		} else { // just varName
			// helpWrite(X.file, X.tokenizer.FormatTokenString()) // write the varName
			helpWrite(X.xmlFile, X.symbolTable.IdentifierToXML(X.tokenizer.Token, false))
			varname := X.tokenizer.Token
			X.vmwriter.WritePush(X.symbolTable.KindOf(varname), X.symbolTable.IndexOf(varname)) //push variable
			X.tokenizer.Advance()
		}

	} else if firstType == "SYMBOL" { // '(' expression ')' | unaryOp term
		if X.tokenizer.Token == "(" { // '(' expression ')'
			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the '('
			X.tokenizer.Advance()
			X.CompileExpression()
			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the ')'
			X.tokenizer.Advance()

		} else { // unaryOp term
			helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the unaryOp
			unary := X.tokenizer.Token
			unary_type := convert_unary_token(unary)
			X.tokenizer.Advance()
			X.CompileTerm()
			X.vmwriter.WriteArithmetic(unary_type)
		}
	}

	helpWrite(X.xmlFile, "</term>\n")
}

// Compiles a (possibly empty) comma-separated list of expressions.
// Follows grammar: (expression (',' expression)* )?
func (X *CompilationEngine) CompileExpressionList() int {
	// As expressionLists always have a ')' after, check if it does, if so the expression list is empty
	if X.tokenizer.Symbol() == ")" {
		helpWrite(X.xmlFile, "<expressionList> </expressionList>\n")
		// no need to advacnce, the ') is part of the caller of expressionList
		return 0
	}
	helpWrite(X.xmlFile, "<expressionList>\n")
	i := 1
	// expression
	X.CompileExpression()

	//(',' expression)*
	for X.tokenizer.Symbol() == "," {
		i++
		helpWrite(X.xmlFile, X.tokenizer.FormatTokenString()) // write the comma
		X.tokenizer.Advance()
		X.CompileExpression()
	}

	helpWrite(X.xmlFile, "</expressionList>\n")
	return i
}

func convert_binary_token(token string) string {
	switch token {
	case "+":
		return "ADD"
	case "-":
		return "SUB"
	case ">":
		return "GT"
	case "<":
		return "LT"
	case "=":
		return "EQ"
	case "&":
		return "AND"
	case "|":
		return "OR"
	case "*":
		return "MUL"
	case "/":
		return "DIV"
	default:
		return "ERROR"
	}
}

func convert_unary_token(token string) string {
	switch token {
	case "-":
		return "NEG"
	case "~":
		return "NOT"
	default:
		return "ERROR"
	}
}
