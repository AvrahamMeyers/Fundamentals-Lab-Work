package Parser

import (
	"strings"
)

func Handle_line(line string, counter_pointer *int) string {
	words := strings.Fields(line)
	//todo: check if line is valid

	counter_logical := *counter_pointer
	//convert string to int
	var str_to_add string
	switch words[0] {
	case "add":
		str_to_add = CodeWriter.add()
		// case "sub":
		// 	str_to_add = handleSub()
		// case "neg":
		// 	str_to_add = handleNeg()
		// case "eq":
		// 	counter_logical++
		// 	str_to_add = handleEq(counter_logical)
		// case "gt":
		// 	counter_logical++
		// 	str_to_add = handleGt(counter_logical)
		// case "lt":
		// 	counter_logical++
		// 	str_to_add = handleLt(counter_logical)
		// case "push":
		// 	index, err := strconv.Atoi(words[2])
		// 	if err != nil {
		// 		fmt.Println("Error converting string to integer:", err)
		// 		break
		// 	}
		// 	str_to_add = handlePush(words[1], index)
		// case "pop":
		// 	index, err := strconv.Atoi(words[2])
		// 	if err != nil {
		// 		fmt.Println("Error converting string to integer:", err)
		// 		break
		// 	}
		// 	str_to_add = handlePop(words[1], index)
		// default:
		// 	fmt.Println("Error: not correct line")
	}
	*counter_pointer = counter_logical
	return str_to_add
}
