package otherpackage

import "fmt"

//Write a "Writing from otherpackage" text
func write() {
	fmt.Println("Writing from otherpackage - Private func")
}

//If the func starts with "W" is public, if it's with "w" is private
