package main

import (
	"fmt"
)

type Character struct {
	Name, Class, Race string
}

var (
	gimli = Character{"Gimli","Barbarian","Dwarf"}
	gandalf = Character{"Gandalf","Wizzard","Ainur"}
)

func (c Character) PrintWhoIs() {
	fmt.Printf("%v's class is a %v and his race is %v.\n", c.Name, c.Class, c.Race)
}

func main() {
	gimli.PrintWhoIs()
	gandalf.PrintWhoIs()
	fmt.Printf("The class of %v is %v.\n", gimli.Name, gimli.Class)
	gandalf.Race = "Istari"
	gandalf.PrintWhoIs()
}