package CodeWriter

func Label(funcStr string, label string) string {
	return "(" + funcStr + "$" + label + ")\n"
}
