package main

import (
	"fmt"
)

type Message struct{
  txt string
}

func (m Message) print () {
	fmt.Println(m.txt)
}

func NewMessage(s string) Message {
	return Message{ txt : s }
}

func main() {
	// TODO change the next line to call the constructor
	m := NewMessage("Tra tra")
	m.print()
}
