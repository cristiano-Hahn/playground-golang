package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	error := checkmail.ValidateFormat("cristiano.hahnifood.com.br")
	fmt.Println(error)
}
