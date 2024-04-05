package main

import "fmt"

func main() {
	printType("aritmetics")
	aritmetics()

	printType("atributors")
	atributors()

	printType("logical")
	logical()
}

func aritmetics() {

	fmt.Println(1 + 2)
	fmt.Println(1 - 2)
	fmt.Println(1 * 2)
	fmt.Println(1 / 2.0)
	fmt.Println(1 % 2)

	var num1 int16 = 10
	var num2 int32 = 20
	sum := int32(num1) + num2
	fmt.Println(sum)

	num := 10
	num++
	fmt.Println(num)
	num += 1
	fmt.Println(num)
	num--
	fmt.Println(num)
	num -= 1
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 3
	fmt.Println(num)
}

func atributors() {
	var var1 string = "string"
	var2 := "asd"

	fmt.Println(var1)
	fmt.Println(var2)
}

func logical() {
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func printType(text string) {
	fmt.Println("")
	fmt.Println(text + "-> ")
	fmt.Println("------------------")
}
