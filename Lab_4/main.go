package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_4/CompilationEngine"
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

func test_tokenizer(filepath string, folderpath string) {
	// filepath := "test.jack"

	// folderpath := "C:\\Users\\csfwn\\school\\Fundamentals\\Fundamentals-Lab-Work\\Fundamentals-Lab-Work\\Lab_4\\Tokenizer"

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

func readFolderPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the folder path: ")
	folderPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "Error"
	}
	// Remove the newline character from the end of the input
	folderPath = strings.TrimSpace(folderPath)
	return folderPath
}

func readFileNamesFromFolder(folder_name string) []string {
	//read all file names in the folder

	file_name_list := []string{}
	file_names, err := os.ReadDir(folder_name)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return file_name_list
	}
	for _, file := range file_names {
		if file.Type().IsRegular() {
			if strings.Split(file.Name(), ".")[1] == "jack" {
				file_name_list = append(file_name_list, file.Name())
			}
		}
	}
	return file_name_list
}

func compilation_main() {
	folder_path := readFolderPath()
	file_names := readFileNamesFromFolder(folder_path)

	for _, file_name := range file_names {

		file_title := strings.Split(file_name, ".")[0]

		var compiler CompilationEngine.CompilationEngine
		compiler.Constructor(file_title, folder_path)

		fmt.Println("End of input file: ", file_title)
		// test_tokenizer(file_title, folder_path)
	}
}

func main() {
	compilation_main()
}
