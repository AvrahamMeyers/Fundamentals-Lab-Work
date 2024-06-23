package main

import (
	"fmt"
	"os"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

func print_tokenizer_token_info(tokenizer Tokenizer.Tokenizer) {
	fmt.Print(tokenizer.Token, " ", tokenizer.TokenType(), " ")

	if tokenizer.TokenType() == "KEYWORD" {
		fmt.Println(tokenizer.KeyWord())
	}
	if tokenizer.TokenType() == "IDENTIFIER" {
		fmt.Println(tokenizer.Identifier())
	}
	if tokenizer.TokenType() == "SYMBOL" {
		fmt.Println(tokenizer.Symbol())
	}
	if tokenizer.TokenType() == "INT_CONST" {
		fmt.Println(tokenizer.IntVal())
	}
	if tokenizer.TokenType() == "STRING_CONST" {
		fmt.Println(tokenizer.StringVal())
	}

	fmt.Printf("File Pos: %d\n", tokenizer.FilePos)
	fmt.Println()
}

func test_tokenizer() {
	filepath := "test.jack"

	folderpath := "C:\\Users\\csfwn\\school\\Fundamentals\\Fundamentals-Lab-Work\\Fundamentals-Lab-Work\\Lab_4\\Tokenizer"

	var tokenizer Tokenizer.Tokenizer
	tokenizer.Constructor(filepath, folderpath)
	// tokenizer.Advance()

	// for tokenizer.HasMoreTokens() {
	// 	print_tokenizer_token_info(tokenizer)
	// 	tokenizer.Advance()
	// }
	// print_tokenizer_token_info(tokenizer)

	outputFile, err := os.Create("output.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	final_str := tokenizer.TokenizeFile()
	_, err = outputFile.WriteString(final_str)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
}

func main() {
	test_tokenizer()
}
