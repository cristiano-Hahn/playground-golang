package main

import "fmt"

func main() {
	var varString string = "String var explicit"
	fmt.Println(varString)

	separator()

	varString2 := "String var infered"
	fmt.Println(varString2)

	separator()

	varString3, varString4 := "String var 3", "String var 4"
	fmt.Println(varString3)
	fmt.Println(varString4)

	separator()

	varString3, varString4 = varString4, varString3
	fmt.Println(varString3)
	fmt.Println(varString4)

	separator()

	const constant = "Constant"
	fmt.Println(constant)

	separator()
}

func separator() {
	fmt.Println("-----------------------------")
}
