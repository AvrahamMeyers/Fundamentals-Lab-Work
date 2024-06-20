package Tokenizer

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func removeComments(text string) string {
	// Compile the regular expression for block comments
	reBlock, err := regexp.Compile(`/\*[\s\S]*?\*/`)
	if err != nil {
		fmt.Println("Error compiling block comment regex:", err)
		return text
	}

	// Remove block comments
	text = reBlock.ReplaceAllString(text, "")

	// Compile the regular expression for inline comments
	reLine, err := regexp.Compile(`(?m)//.*$`)
	if err != nil {
		fmt.Println("Error compiling line comment regex:", err)
		return text
	}

	// Remove inline comments
	text = reLine.ReplaceAllString(text, " ")

	// Remove any trailing whitespace from each line
	text = regexp.MustCompile(`(?m)^\s+|\s+$`).ReplaceAllString(text, "")

	// Remove any extra empty lines left behind
	text = regexp.MustCompile(`(?m)^\s*\n`).ReplaceAllString(text, "")

	// remove all white space
	//text = regexp.MustCompile(`\s*`).ReplaceAllString(text, "")

	return text
}

func inList(text string, list []string) bool {
	for _, keyword := range list {
		if text == keyword {
			return true
		}
	}
	return false
}

func isKeyWord(text string) bool {
	/*Check if the provided word is equal to one of the keywords.*/
	keywords := []string{"class", "constructor", "function", "method", "field", "static", "var",
		"int", "char", "boolean", "void", "true", "false", "null", "this",
		"let", "do", "if", "else", "while", "return"}
	return inList(text, keywords)
}

func isSymbol(text string) bool {
	symbols := []string{"{", "}", "(", ")", "[", "]",
		".", ",", ";",
		"+", "-", "*", "/", "&",
		"|", "<", ">", "=", "~"}
	return inList(text, symbols)
}

func isIdentifier(text string) bool {
	re, err := regexp.Compile(`^[A-Za-z_][A-Za-z0-9_]*$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	if re.MatchString(text) {
		return true
	}
	return false
}

func isInt_Const(text string) bool {
	// Compile the regular expression for a valid integer
	re, err := regexp.Compile(`^[0-9]+$`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	// Check if the text matches the regular expression
	if !re.MatchString(text) {
		return false
	}

	// Convert the string to an integer
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return false
	}

	// Check if the integer is within the specified range
	return num >= 0 && num <= 32767
}

func isString_Const(text string) bool {
	// Compile the regular expression
	re, err := regexp.Compile(`"[^\"\n]*"`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	if re.MatchString(text) {
		return true
	}
	return false
}

func convertSymbolForXML(symbol string) string {
	symbolsToConvert := []string{">", "<", "\"", "&"}

	if inList(symbol, symbolsToConvert) {
		if symbol == ">" {
			return "&gt;"
		}
		if symbol == "<" {
			return "&lt;"
		}
		if symbol == "\"" {
			return "&quot;"
		}
		if symbol == "&" {
			return "&amp;"
		}
	}
	return symbol
}

type Tokenizer struct {
	Token    string
	Filetext string
	//as an example
	//fp os.File
	//here we may want items such as current token file pointer
}

func (X *Tokenizer) Constructor(fileName string, folderpath string) {
	/*Opens the input file/stream and gets
	ready to tokenize it.*/
	file_path := folderpath + "/" + fileName
	file, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("Error opening file: ", fileName, err)
		return
	}
	X.Filetext = removeComments(string(file))
	X.Token = ""
}

// Do we have more tokens in the input?
func (X *Tokenizer) HasMoreTokens() bool {
	return false
}

// Gets the next token from the input
// and makes it the current token. This
// method should only be called if
// hasMoreTokens() is true. Initially
// there is no current token.
func (X *Tokenizer) Advance() {
	if X.HasMoreTokens() {

	}

}

// Returns the type of the current token.
// Possible return values:
// KEYWORD, SYMBOL, IDENTIFIER, INT_CONST, STRING_CONST
func (X *Tokenizer) TokenType() string {
	if isKeyWord(X.Token) {
		return "KEYWORD"
	}
	if isSymbol(X.Token) {
		return "SYMBOL"
	}
	if isIdentifier(X.Token) {
		return "IDENTIFIER"
	}
	if isInt_Const(X.Token) {
		return "INT_CONST"
	}
	if isString_Const(X.Token) {
		return "STRING_CONST"
	}
	return ""
}

// Returns the keyword which is the current token. Should be called only
// when tokenType() is KEYWORD. Possible return values:
// CLASS, METHOD, FUNCTION, CONSTRUCTOR, INT, BOOLEAN, CHAR, VOID,
// VAR, STATIC, FIELD, LET, DO, IF, ELSE, WHILE,
// RETURN, TRUE, FALSE, NULL, THIS
// e.g. <keyword> CLASS </keyword>
func (X Tokenizer) KeyWord() string {
	return strings.ToUpper(X.Token)

	/*
		if isKeyWord(X.Token) {
			return strings.ToUpper(X.Token)
		}
		return ""
	*/
}

// Returns the character which is the
// current token. Should be called only
// when tokenType() is SYMBOL. e.g. <symbol> + </symbol>
func (X *Tokenizer) Symbol() string {
	return convertSymbolForXML(X.Token)

	/*
		if isSymbol(X.Token) {
			return convertSymbolForXML(X.Token)
		}
		return ""
	*/
}

// Returns the identifier which is the
// current token. Should be called only
// when tokenType() is IDENTIFIER. e.g. <identifier> varName </identifier>
func (X *Tokenizer) Identifier() string {
	return X.Token
}

// Returns the integer value of the
// current token. Should be called only
// when tokenType() is INT_CONST. e.g. <integerConstant> 33 </integerConstant>
func (X *Tokenizer) IntVal() int {
	num, err := strconv.Atoi(X.Token)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return num
}

// Returns the string value of the current
// token, without the double quotes.
// Should be called only when
// tokenType() is STRING_CONST. e.g. <stringConstant> foo </stringConstant>
func (X *Tokenizer) StringVal() string {
	return ""
}
