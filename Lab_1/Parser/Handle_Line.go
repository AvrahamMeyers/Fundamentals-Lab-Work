package Parser

import (
	"fmt"
	"strings"

	"github.com/AvrahamMeyers/Fundamentals-Lab-Work/Lab_1/CodeWriter"
)

// Handle_line takes a line of code and returns the corresponding assembly code
// scope is a pointer to a string that is used to keep track of which scope that line is in: Global, or a spefic function
func Handle_line(line string, fileName string, counter int, scope *string) string {
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
		str_to_add = CodeWriter.Eq(counter)
		counter += 1
	case "gt":
		str_to_add = CodeWriter.Gt(counter)
		counter += 1
	case "lt":
		str_to_add = CodeWriter.Lt(counter)
		counter += 1
	case "and":
		str_to_add = CodeWriter.And()
	case "or":
		str_to_add = CodeWriter.Or()
	case "not":
		str_to_add = CodeWriter.Not()
	case "push":
		var str = handlePush(words[1], words[2], fileName)
		if str == "Error" {
			fmt.Println("Error: not correct line")
			break
		}
		str_to_add = str
	case "pop":
		var str = handlePop(words[1], words[2], fileName)
		if str == "Error" {
			fmt.Println("Error: not correct line")
			break
		}
		str_to_add = str
	case "//":
		str_to_add = line + "\n"
	case "label":
		str_to_add = CodeWriter.Label(*scope, words[1])
	case "if-goto":
		str_to_add = CodeWriter.GotoIf(*scope, words[1])
	case "goto":
		str_to_add = CodeWriter.Goto(*scope, words[1])
	case "function":
		// adding a comment to the file.
		str_to_add += CodeWriter.Function(words[1], words[2])

		//function format: a function is declared function=(words[0])
		//function.name=(wordss[1]),
		*scope = words[1]
	case "call":
		str_to_add = CodeWriter.Call(words[1], words[2], counter)
		counter += 1
	case "return":
		str_to_add = CodeWriter.Return()
	default:
		str_to_add = "//The following line was not handled: " + line + "\n"
	}
	return "//" + line + "\n" + str_to_add

}

func handlePush(segment string, index string, fileName string) string {
	var str_to_add string
	switch segment {
	case "argument":
		str_to_add = CodeWriter.PushArgument(index)
	case "local":
		str_to_add = CodeWriter.PushLocal(index)
	case "static":
		str_to_add = CodeWriter.PushStatic(index, fileName)
	case "constant":
		str_to_add = CodeWriter.PushConstant(index)
	case "this":
		str_to_add = CodeWriter.PushThis(index)
	case "that":
		str_to_add = CodeWriter.PushThat(index)
	case "temp":
		str_to_add = CodeWriter.PushTemp(index)
	case "pointer":
		str_to_add = CodeWriter.PushPointer(index)

	default:
		str_to_add = "Error"
	}
	return str_to_add
}

func handlePop(segment string, index string, fileName string) string {
	var str_to_add string
	switch segment {
	case "argument":
		str_to_add = CodeWriter.PopArgument(index)
	case "local":
		str_to_add = CodeWriter.PopLocal(index)
	case "static":
		str_to_add = CodeWriter.PopStatic(index, fileName)
	case "constant":
		str_to_add = "Error"
	case "this":
		str_to_add = CodeWriter.PopThis(index)
	case "that":
		str_to_add = CodeWriter.PopThat(index)
	case "temp":
		str_to_add = CodeWriter.PopTemp(index)
	case "pointer":
		str_to_add = CodeWriter.PopPointer(index)
	default:
		str_to_add = "Error"
	}
	return str_to_add
}
