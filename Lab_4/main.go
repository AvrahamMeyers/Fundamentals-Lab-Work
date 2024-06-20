package main

import (
	"fmt"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/Tokenizer"
)

func test_tokenizer() {
	filepath := "test.jack"

	folderpath := "C:\\Users\\a3210\\Machon Lev\\Year 4\\Semester B\\Fundamentals\\Project\\Fundamentals-Lab-Work\\Lab_4\\Tokenizer"

	var tokenizer Tokenizer.Tokenizer
	tokenizer.Constructor(filepath, folderpath)

	fmt.Println(tokenizer.Filetext)
}

func main() {
	test_tokenizer()
}
