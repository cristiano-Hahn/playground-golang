package otherpackage

import "fmt"

//Write a "Writing from otherpackage" text
func Write() {
	fmt.Println("Writing from otherpackage")
	write()
}

//If the func starts with "W" is public, if it's with "w" is private
