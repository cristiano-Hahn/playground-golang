package main

import "fmt"

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	street string
	number int8
}

func main() {
	var user1 user
	user1.age = 18
	user1.name = "cris"
	user1.address.street = "street"
	user1.address.number = 1
	fmt.Println(user1)

	user2 := user{"cris", 18, address{"street", 1}}
	fmt.Println(user2)

	user3 := user{name: "cris"}
	fmt.Println(user3)

}
