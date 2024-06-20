package Tokenizer

//import "os"

type Tokenizer struct {
	Token string
	//as an example
	//fp os.File
	//here we may want items such as current token file pointer
}

func (X *Tokenizer) Constructor(fileName string) {
	/*Opens the input file/stream and gets
	ready to tokenize it.*/

}

//Do we have more tokens in the input?
func (X *Tokenizer) HasMoreTokens() bool {
	return false
}

//Gets the next token from the input
//and makes it the current token. This
//method should only be called if
//hasMoreTokens() is true. Initially
//there is no current token.
func (X *Tokenizer) Advance() {

}

// Returns the type of the current token.
// Possible return values:
// KEYWORD, SYMBOL, IDENTIFIER, INT_CONST, STRING_CONST
func (X *Tokenizer) TokenType() string {

	return ""
}

// Returns the keyword which is the current token. Should be called only
// when tokenType() is KEYWORD.Possible return values:
// CLASS, METHOD, FUNCTION, CONSTRUCTOR, INT, BOOLEAN, CHAR, VOID,
// VAR, STATIC, FIELD, LET, DO, IF, ELSE, WHILE,
// RETURN, TRUE, FALSE, NULL, THIS
func (X Tokenizer) KeyWord() string {

	return ""
}

//Returns the character which is the
//current token. Should be called only
//when tokenType() is SYMBOL.
func (X *Tokenizer) Symbol() string {

	return ""
}

//Returns the identifier which is the
//current token. Should be called only
//when tokenType() is IDENTIFIER.
func (X *Tokenizer) Identifier() {

}

//Returns the integer value of the
//current token. Should be called only
//when tokenType() is INT_CONST.
func (X *Tokenizer) IntVal() {

}

// Returns the string value of the current
//token, without the double quotes.
//Should be called only when
//tokenType() is STRING_CONST.
func (X *Tokenizer) StringVal() {

}
