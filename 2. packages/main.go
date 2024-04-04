package main

import (
	"fmt"
	"module/otherpackage"
)

func main() {
	fmt.Println("Writing from main file")
	otherpackage.Write()
	fmt.Scanln()
}

//run "go build" to generate binary files
