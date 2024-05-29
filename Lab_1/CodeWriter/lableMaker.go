package CodeWriter

func Label(fnction string, lbl string) string {
	return "(" + fnction + "$" + lbl + ")\n"
}
