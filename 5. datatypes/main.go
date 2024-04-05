package main

import (
	"errors"
	"fmt"
)

func main() {
	printType("Integers")
	integers()

	printType("Floats")
	floats()

	printType("Strings")
	strings()

	printType("Booleans")
	booleans()

	printType("Errors")
	err()
}

func integers() {
	//can be int8, int16, int32, int64 or just int, that consider de processor architecture
	var integer int16 = 100
	fmt.Println(integer)

	//uint is an int without consider the signal
	var integer2 uint16 = 100
	fmt.Println(integer2)

	//rune is an alias for int32, commonly used to represent ascii characters
	var integer3 rune = 100
	fmt.Println(integer3)

	//byte is a uint8 alias
	var integer4 byte = 100
	fmt.Println(integer4)
}

func floats() {
	//can be float32, float64 or just float, that consider de processor architecture
	var float float32 = 123.45
	fmt.Println(float)
}

func strings() {
	var string string = "text"
	fmt.Println(string)

	var char rune = 'A'
	fmt.Println(char)
}

func booleans() {
	var boolean bool = false
	fmt.Println(boolean)
}

func err() {
	var err error = nil
	fmt.Println(err)

	var err2 error = errors.New("internal error")
	fmt.Println(err2)
}

func printType(text string) {
	fmt.Println("------------------")
	fmt.Println(text + "-> ")
}
