package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tar0 string = "tar0"

/*Should print to the end of output file the following:

command: add
or
command: sub
or
command: neg
*/

func handleAdd() string {
	return "command: add\n"
}

func handleSub() string {
	return "command: sub\n"
}

func handleNeg() string {
	return "command: neg\n"
}

/*End Arithmetic*/

/*Logic:
Should print to the end of output file the following:

command: eq
or
command: gt
or
command: lt

After that it should print the current counter value and increment the counter. For example

counter: 3
*/

func handleEq(cl int) string {
	return "command: eq counter: " + strconv.Itoa(cl) + "\n"
}

func handleGt(cl int) string {
	return "command: gt counter: " + strconv.Itoa(cl) + "\n"
}

func handleLt(cl int) string {
	return "command: lt counter: " + strconv.Itoa(cl) + "\n"
}

/*END LOGIC*/
/*
Memory access:
Should print to the end of output file the following:

command: push segment: <s> index: <i>
or
command: pop segment: <s> index: <i>

For example, for the input
push static 2

We should see in the file
command: push segment: static index: 2

*/
func handlePush(segment string, index int) string {
	return "command: push segment: " + segment + " index: " + strconv.Itoa(index) + "\n"
}

func handlePop(segment string, index int) string {
	return "command: pop segment: " + segment + " index: " + strconv.Itoa(index) + "\n"
}

/*END MEMORY ACCESS*/

func readfile() {
	var counter_logical int = 0
	//open and read the file
	data, err := os.ReadFile("inputA.vm")
	if err != nil {
		fmt.Print(err)
	}
	file, err := os.OpenFile(tar0+".asm", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	//fmt.Println(string(data))
	datastring := string(data)

	//split the text into lines
	lines := strings.Split(datastring, "\n")

	for _, line := range lines {
		//split the line into an array(slice)
		words := strings.Fields(line)
		//todo: check if line is valid

		//convert string to int
		var str_to_add string
		switch words[0] {
		case "add":
			str_to_add = handleAdd()
		case "sub":
			str_to_add = handleSub()
		case "neg":
			str_to_add = handleNeg()
		case "eq":
			counter_logical++
			str_to_add = handleEq(counter_logical)
		case "gt":
			counter_logical++
			str_to_add = handleGt(counter_logical)
		case "lt":
			counter_logical++
			str_to_add = handleLt(counter_logical)
		case "push":
			index, err := strconv.Atoi(words[2])
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				break
			}
			str_to_add = handlePush(words[1], index)
		case "pop":
			index, err := strconv.Atoi(words[2])
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				break
			}
			str_to_add = handlePush(words[1], index)
		default:
			fmt.Println("Error: not correct line")
		}
		_, err = file.WriteString(str_to_add)
		if err != nil {
			fmt.Println("Error appending to file:", err)
			return
		}
	}

}

func main() {

	readfile()
}
