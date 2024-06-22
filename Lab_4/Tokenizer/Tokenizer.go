package Tokenizer

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Tokenizer struct {
	Token    string
	Filetext string
	FilePos  int
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
	X.FilePos = 0
}

// Do we have more tokens in the input?
func (X *Tokenizer) HasMoreTokens() bool {
	return X.FilePos < len(X.Filetext)
}

// skipWhitespace skips whitespace characters in the input
func (t *Tokenizer) skipWhitespace() {
	for t.FilePos < len(t.Filetext) && unicode.IsSpace(rune(t.Filetext[t.FilePos])) {
		t.FilePos++
	}
}

// Gets the next token from the input
// and makes it the current token. This
// method should only be called if
// hasMoreTokens() is true. Initially
// there is no current token.
// Advance moves to the next token and updates the current token
func (t *Tokenizer) Advance() {
	t.skipWhitespace()
	if !t.HasMoreTokens() {
		return
	}

	var tokenValue strings.Builder
	ch := string(t.Filetext[t.FilePos])

	// Handle symbols
	if isSymbol(ch) {
		t.Token = ch
		t.FilePos++
		return
	}

	// Handle string constants
	runeCh := rune(ch[0])
	if runeCh == '"' {
		t.FilePos++
		for t.FilePos < len(t.Filetext) && rune(t.Filetext[t.FilePos]) != '"' {
			tokenValue.WriteByte(t.Filetext[t.FilePos])
			t.FilePos++
		}
		t.Token = tokenValue.String()
		t.FilePos++ // Skip closing quote
		return
	}

	// Handle integer constants
	if unicode.IsDigit(runeCh) {
		for t.FilePos < len(t.Filetext) && unicode.IsDigit(rune(t.Filetext[t.FilePos])) {
			tokenValue.WriteByte(t.Filetext[t.FilePos])
			t.FilePos++
		}
		t.Token = tokenValue.String()
		return
	}

	// Handle identifiers and keywords
	if unicode.IsLetter(runeCh) || runeCh == '_' {
		for t.FilePos < len(t.Filetext) &&
			(unicode.IsLetter(rune(t.Filetext[t.FilePos])) ||
				unicode.IsDigit(rune(t.Filetext[t.FilePos])) ||
				rune(t.Filetext[t.FilePos]) == '_') {
			tokenValue.WriteByte(t.Filetext[t.FilePos])
			t.FilePos++
		}
		value := tokenValue.String()
		t.Token = value
		return
	}

	// Unexpected character
	fmt.Printf("Unexpected character: %c\n", runeCh)
	t.FilePos++
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
func (X Tokenizer) KeyWord() string {
	return strings.ToLower(X.Token)
}

// Returns the character which is the
// current token. Should be called only
// when tokenType() is SYMBOL.
func (X *Tokenizer) Symbol() string {
	return convertSymbolForXML(X.Token)
}

// Returns the identifier which is the
// current token. Should be called only
// when tokenType() is IDENTIFIER.
func (X *Tokenizer) Identifier() string {
	return X.Token
}

// Returns the integer value of the
// current token. Should be called only
// when tokenType() is INT_CONST.
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
	return removeQuotes(X.Token)
}

// Returns the XML format of a token of the following types:
// <keyword> CLASS </keyword>
// <identifier> varName </identifier>
// <symbol> + </symbol>
// <integerConstant> 33 </integerConstant>
// <stringConstant> foo </stringConstant>
func (X *Tokenizer) FormatTokenString() string {
	if X.TokenType() == "KEYWORD" {
		return "<keyword> " + X.KeyWord() + " </keyword>\n"
	}
	if X.TokenType() == "IDENTIFIER" {
		return "<identifier> " + X.Identifier() + " </identifier>\n"
	}
	if X.TokenType() == "SYMBOL" {
		return "<symbol> " + X.Symbol() + " </symbol>\n"
	}
	if X.TokenType() == "INT_CONST" {
		return "<integerConstant> " + fmt.Sprint(X.IntVal()) + " </integerConstant>\n"
	}
	if X.TokenType() == "STRING_CONST" {
		return "<stringConstant> " + X.StringVal() + " </stringConstant>\n"
	}
	return "invalid token"
}

func (X *Tokenizer) TokenizeFile() string {
	final_str := "<tokens>\n"

	for X.HasMoreTokens() {
		final_str += X.FormatTokenString()
		X.Advance()
	}
	final_str += X.FormatTokenString()
	final_str += "</tokens>"

	return final_str
}

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
	keywords := []string{"class", "Class", "constructor", "function", "method", "field", "static", "var",
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
	re, err := regexp.Compile(`[^\"\n]*`)
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

func removeQuotes(s string) string {
	// Trim leading and trailing double quotes
	return strings.Trim(s, `"`)
}
