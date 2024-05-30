package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_1/CodeWriter"
	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_1/Parser"
)

func processFile(input_file_content string, output_file *os.File, fileName string) {

	//fmt.Println(string(data))
	datastring := input_file_content

	//split the text into lines
	lines := strings.Split(datastring, "\n")
	counter := 0

	// scope is a string that is used to keep track of which scope that line is in: Global, or a spefic function
	scope := "Global"

	for _, line := range lines {
		words := strings.Fields(line)
		if len(words) > 0 {
			//split the line into an array(slice)
			str_to_add := Parser.Handle_line(line, fileName, counter, &scope)
			counter++
			var err error
			_, err = output_file.WriteString(str_to_add)
			if err != nil {
				fmt.Println("Error appending to file:", err)
				return
			}
		}
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
			if strings.Split(file.Name(), ".")[1] == "vm" {
				file_name_list = append(file_name_list, file.Name())
			}
		}
	}
	return file_name_list
}

func new_main() {
	folder_path := readFolderPath()
	file_names := readFileNamesFromFolder(folder_path)

	// Get the base name of the path
	folder_name := filepath.Base(folder_path)

	output_file, err := os.Create(folder_path + "/" + folder_name + ".asm")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	_, err = output_file.WriteString(CodeWriter.Bootsrap())
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	for _, file_name := range file_names {

		file_title := strings.Split(file_name, ".")[0]

		file_path := folder_path + "/" + file_name
		input_file, err := os.ReadFile(file_path)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		input_file_content := string(input_file)

		processFile(input_file_content, output_file, file_title)

		fmt.Println("End of input file: ", file_title)
	}
}

func main() {
	new_main()
}
