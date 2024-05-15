package main

func add() string {
	return "@sp\n" +
		"D = M\n" +
		"A = A - 1\n" +
		"MD = D + M\n"
}

func sub() string {
	return "@x \n" +
		"D = A\n" +
		"@y \n" +
		"D = D - A\n"
}

func neg() string {
	return "@y \n" +
		"D = -A\n"
}
