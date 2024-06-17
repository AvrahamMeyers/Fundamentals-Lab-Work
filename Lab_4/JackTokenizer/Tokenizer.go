package jacktokenizer

func Constructor() {
	/*Opens the input file/stream and gets
	ready to tokenize it.*/
}

func HasMoreTokens() bool {
	//Do we have more tokens in the input?
	return false
}

func Advance() {
	/*Gets the next token from the input
	and makes it the current token. This
	method should only be called if
	hasMoreTokens() is true. Initially
	there is no current token.*/
}

func TokenType() {
	//Returns the type of the current token
}

func keyWord() {
	/*Returns the keyword which is the
	current token. Should be called only
	when tokenType() is KEYWORD.*/
}

func Symbol() {
	/*Returns the character which is the
	current token. Should be called only
	when tokenType() is SYMBOL.*/
}

func Identifier() {
	/*Returns the identifier which is the
	current token. Should be called only
	when tokenType() is IDENTIFIER.*/
}

func IntVal() {
	/*Returns the integer value of the
	current token. Should be called only
	when tokenType() is INT_CONST.*/
}

func StringVal() {
	/* Returns the string value of the current
	token, without the double quotes.
	Should be called only when
	tokenType() is STRING_CONST.*/
}
