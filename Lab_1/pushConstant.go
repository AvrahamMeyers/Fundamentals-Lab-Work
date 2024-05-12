package main

//expected is just the number string return the commands @num  D=A
func pushConstant(num string) string {
	return "@" + num + "\nD = A"
}
