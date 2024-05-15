package Parser

import (
	"fmt"
	"strings"

	"../CodeWriter"
)

func Handle_line(line string) string {
	words := strings.Fields(line)
	//todo: check if line is valid

	var str_to_add string
	switch words[0] {
	case "add":
		str_to_add = CodeWriter.Add()
	case "sub":
		str_to_add = CodeWriter.Sub()
	case "neg":
		str_to_add = CodeWriter.Neg()
	case "eq":
		str_to_add = CodeWriter.Eq()
	case "gt":
		str_to_add = CodeWriter.Gt()
	case "lt":
		str_to_add = CodeWriter.Lt()
	case "and":
		str_to_add = CodeWriter.And()
	case "or":
		str_to_add = CodeWriter.Or()
	case "not":
		str_to_add = CodeWriter.Not()
	case "push":
		var str = handlePush(words[1], words[2])
		if str == "Error" {
			fmt.Println("Error: not correct line")
			break
		}
		str_to_add = str
	case "pop":
		var str = handlePop(words[1], words[2])
		if str == "Error" {
			fmt.Println("Error: not correct line")
			break
		}
		str_to_add = str
	default:
		fmt.Println("Error: not correct line")
	}
	return str_to_add
}

func handlePush(segment string, index string) string {
	var str_to_add string
	switch segment {
	// case "argument":
	// 	str_to_add = CodeWriter.PushArgument(index)
	// case "local":
	// 	str_to_add = CodeWriter.PushLocal(index)
	// case "static":
	// 	str_to_add = CodeWriter.PushStatic(index)
	case "constant":
		str_to_add = CodeWriter.PushConstant(index)

	// case "this":
	// 	str_to_add = CodeWriter.PushThis(index)
	// case "that":
	// 	str_to_add = CodeWriter.PushThat(index)
	// case "temp":
	// 	str_to_add = CodeWriter.PushTemp(index)
	// case "pointer":
	// 	str_to_add = CodeWriter.PushPointer(index)

	default:
		str_to_add = "Error"
	}
	return str_to_add
}

func handlePop(segment string, index string) string {
	var str_to_add string
	switch segment {
	// case "argument":
	// 	str_to_add = CodeWriter.PopArgument(index)
	// case "local":
	// 	str_to_add = CodeWriter.PopLocal(index)
	// case "static":
	// 	str_to_add = CodeWriter.PopStatic(index)
	// case "constant":
	// 	str_to_add = CodeWriter.PopConstant(index)
	// case "this":
	// 	str_to_add = CodeWriter.PopThis(index)
	// case "that":
	// 	str_to_add = CodeWriter.PopThat(index)
	// case "temp":
	// 	str_to_add = CodeWriter.PopTemp(index)
	// case "pointer":
	// 	str_to_add = CodeWriter.PopPointer(index)
	default:
		str_to_add = "Error"
	}
	return str_to_add
}
