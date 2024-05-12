package main

//expected is just the number string return the commands @num  D=A
func pushConstant(num string) string {
	return "D=A\n@" + num //at first I did @num\nD = A but logically this didn't make sense because when we push another constant we need be able to keep both
}
