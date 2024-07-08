package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_5/CompilationEngine"
)

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
	}
}

func main() {
	compilation_main()
}
