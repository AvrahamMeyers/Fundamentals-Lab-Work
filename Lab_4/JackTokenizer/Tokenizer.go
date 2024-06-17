package jacktokenizer

type Tokenizer struct {
	//here we may want items such as current token file pointer
}

func (Tokenizer) Constructor() {
	/*Opens the input file/stream and gets
	ready to tokenize it.*/
}

func (Tokenizer) HasMoreTokens() bool {
	//Do we have more tokens in the input?
	return false
}

func (Tokenizer) Advance() {
	/*Gets the next token from the input
	and makes it the current token. This
	method should only be called if
	hasMoreTokens() is true. Initially
	there is no current token.*/
}

func (Tokenizer) TokenType() {
	//Returns the type of the current token
}

func (Tokenizer) keyWord() {
	/*Returns the keyword which is the
	current token. Should be called only
	when tokenType() is KEYWORD.*/
}

func (Tokenizer) Symbol() {
	/*Returns the character which is the
	current token. Should be called only
	when tokenType() is SYMBOL.*/
}

func (Tokenizer) Identifier() {
	/*Returns the identifier which is the
	current token. Should be called only
	when tokenType() is IDENTIFIER.*/
}

func (Tokenizer) IntVal() {
	/*Returns the integer value of the
	current token. Should be called only
	when tokenType() is INT_CONST.*/
}

func (Tokenizer) StringVal() {
	/* Returns the string value of the current
	token, without the double quotes.
	Should be called only when
	tokenType() is STRING_CONST.*/
}
