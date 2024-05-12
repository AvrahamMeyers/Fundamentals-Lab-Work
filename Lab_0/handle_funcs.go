package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Test_handle_line() {

	var counter_logical int = 0

	line := "add"
	str_to_add := Handle_line(line, &counter_logical)
	if str_to_add != "command: add\n" {
		fmt.Println("Error: add")
	}
	line = "sub"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: sub\n" {
		fmt.Println("Error: sub")
	}
	line = "neg"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: neg\n" {
		fmt.Println("Error: neg")
	}
	line = "eq"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: eq counter: 1\n" {
		fmt.Println("Error: eq")
	}
	line = "gt"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: gt counter: 2\n" {
		fmt.Println("Error: gt got:", str_to_add)
	}
	line = "lt"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: lt counter: 3\n" {
		fmt.Println("Error: lt got:", str_to_add)
	}
	line = "push constant 2"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: push segment: constant index: 2\n" {
		fmt.Println("Error: push constant 2")
	}
	line = "pop constant 2"
	str_to_add = Handle_line(line, &counter_logical)
	if str_to_add != "command: pop segment: constant index: 2\n" {
		fmt.Println("Error: pop constant 2\nexpected: pop constant 2\ngot:", str_to_add)
	}
}

func Handle_line(line string, counter_pointer *int) string {
	words := strings.Fields(line)
	//todo: check if line is valid

	counter_logical := *counter_pointer
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
		str_to_add = handlePop(words[1], index)
	default:
		fmt.Println("Error: not correct line")
	}
	*counter_pointer = counter_logical
	return str_to_add
}

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
