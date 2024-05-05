package main

import (
	"fmt"
	"os"
	"strings"
)

var tar0 string = "tar0"

func processFile(input_file_content string, output_file *os.File) {
	var counter_logical int = 0

	//fmt.Println(string(data))
	datastring := input_file_content

	//split the text into lines
	lines := strings.Split(datastring, "\n")

	for _, line := range lines {
		//split the line into an array(slice)
		str_to_add := Handle_line(line, &counter_logical)

		var err error
		_, err = output_file.WriteString(str_to_add)
		if err != nil {
			fmt.Println("Error appending to file:", err)
			return
		}
	}
}

func readFolderName() string {
	var input string
	fmt.Scanln(&input)
	return input
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
			file_name_list = append(file_name_list, file.Name())
		}
	}
	return file_name_list
}

func new_main() {
	folder_name := readFolderName()
	file_names := readFileNamesFromFolder(folder_name)

	output_file, err := os.Create(tar0 + ".asm")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	for _, file_name := range file_names {
		file_path := folder_name + "/" + file_name
		input_file, err := os.ReadFile(file_path)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		file_title := strings.Split(file_name, ".")[0]

		input_file_content := string(input_file)

		processFile(input_file_content, output_file)

		fmt.Println("End of input file: ", file_title)
	}
}

func test_main() {
	Test_handle_line()

}

func main() {
	test_main()
}
