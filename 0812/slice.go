package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["1"] = PersonInfo{"1", "turtle", "A102..."}
	personDB["2"] = PersonInfo{"2", "two", "B417..."}

	person, ok := personDB["1"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1")
	} else {
		fmt.Println("Did not find person with ID 1")
	}
}
