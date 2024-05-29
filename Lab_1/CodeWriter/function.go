package CodeWriter

import "strconv"

func Function(fname string, lbl string) string {
	str_to_add := "// Now in Function " + fname + "\n(" + fname + ")\n"
	lclVar, err := strconv.Atoi(lbl)
	if err != nil {

	}
	for i := 0; i < lclVar; i++ {
		PushConstant("0")
	}
	return str_to_add
}
