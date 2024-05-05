package main

var tar0 string = "tar0"

// func HandleBuy(ProductName string, Amount int, Price float32) {

// 	//append to tar0.asm file:
// 	//### BUY <ProductName> ###
// 	//<Amount>*<Price>
// }

// func HandleSell(ProductName string, Amount int, Price float32) {

// 	//append to tar0.asm file:
// 	//$$$ CELL <ProductName> $$$
// 	//<Amount>*<Price>

// }

func readfile() {
	//open and read the file
	//data, err := ioutil.ReadFile("inputA.vm")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println(string(data))
	//datastring := string(data)

	//split the text into lines
	//lines := strings.Split(datastring, "\n")

	//for _, line := range lines {
	//split the line into an array(slice)
	//words := strings.Fields(line)
	//check if line is valid
	// if len(words) < 3 { //for some odd reason the last line has [0/0]0x8a8b20 this as part of the line, perhaps an error somewerhe in code
	// 	println(words)
	// 	break
	// }
	//convert string to int
	// amount, err := strconv.Atoi(words[2])
	// if err != nil {
	// 	fmt.Println("Error converting string to integer:", err)
	// 	break
	// }
	// //convert string to float
	// price, err := strconv.ParseFloat(words[3], 32)
	// if err != nil {
	// 	fmt.Println("Error converting string to integer:", err)
	// 	break
	// }
	//enter appropriate function
	// if string(words[0]) == "buy" {
	// 	HandleBuy(string(words[1]), amount, float32(price))
	// } else if string(words[0]) == "cell" {
	// 	HandleSell(string(words[1]), amount, float32(price))
	// }
	//}

}
func main() {
	var counter_logical int = 0

	readfile()
}
