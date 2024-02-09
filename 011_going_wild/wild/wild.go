package main


import (
	"fmt"
  "taylorshaulis/pickword"
  "log"
  )
  

func main() {
  fmt.Println("I am super wild")
  dict := "../dict"
  word, err := pickword.Pick(dict)
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println(word)
  }

}