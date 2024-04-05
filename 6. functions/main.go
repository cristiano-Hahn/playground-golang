package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20))

	resultSum, resultSub := sumAndSub(10, 5)
	fmt.Println(resultSum, resultSub)

	newResultSum, _ := sumAndSub(10, 8)
	fmt.Println(newResultSum)

	fmt.Println(anonymousFunc("Some text"))
}

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func sumAndSub(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func anonymousFunc(text string) string {
	var f = func(text string) string {
		return "The text is " + text
	}
	return f(text)
}
