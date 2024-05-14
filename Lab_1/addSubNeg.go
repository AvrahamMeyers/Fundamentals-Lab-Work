package main

func add() string {
	return "@x \n" +
		"D = A\n" +
		"@y \n" +
		"D = D + A\n"
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
