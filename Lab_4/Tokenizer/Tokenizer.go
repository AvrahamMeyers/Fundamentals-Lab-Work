package Tokenizer

import (
	"fmt"
	"os"
	"regexp"
)

type Tokenizer struct {
	Token    string
	Filetext string
	//as an example
	//fp os.File
	//here we may want items such as current token file pointer
}

func removeCommentsAndWhiteSpace(text string) string {
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

	text = regexp.MustCompile(`\s*`).ReplaceAllString(text, "")

	return text
}

func checkKeyWord(text string) bool {
	return false
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
	X.Filetext = removeCommentsAndWhiteSpace(string(file))
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

	return ""
}

// Returns the keyword which is the current token. Should be called only
// when tokenType() is KEYWORD. Possible return values:
// CLASS, METHOD, FUNCTION, CONSTRUCTOR, INT, BOOLEAN, CHAR, VOID,
// VAR, STATIC, FIELD, LET, DO, IF, ELSE, WHILE,
// RETURN, TRUE, FALSE, NULL, THIS
// e.g. <keyword> CLASS </keyword>
func (X Tokenizer) KeyWord() string {

	return ""
}

// Returns the character which is the
// current token. Should be called only
// when tokenType() is SYMBOL. e.g. <symbol> + </symbol>
func (X *Tokenizer) Symbol() string {

	return ""
}

// Returns the identifier which is the
// current token. Should be called only
// when tokenType() is IDENTIFIER. e.g. <identifier> varName </identifier>
func (X *Tokenizer) Identifier() string {
	return ""
}

// Returns the integer value of the
// current token. Should be called only
// when tokenType() is INT_CONST. e.g. <integerConstant> 33 </integerConstant>
func (X *Tokenizer) IntVal() string {
	return ""
}

// Returns the string value of the current
// token, without the double quotes.
// Should be called only when
// tokenType() is STRING_CONST. e.g. <stringConstant> foo </stringConstant>
func (X *Tokenizer) StringVal() string {
	return ""
}
