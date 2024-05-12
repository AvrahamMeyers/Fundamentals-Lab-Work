package main

// true if x = y and false otherwise

func eq(loc_x string, loc_y string, loc_jump string) string {
	return "@" + loc_x + "\nD = M\n@" + loc_y + "D = D - M\n"
}
