package main


import (
	"fmt"
  "taylorshaulis/pickword"
  )
  

func main() {
  fmt.Println("I am super wild")
  word := pickword.Pick()
  fmt.Println(word)
}