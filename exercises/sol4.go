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
	m := NewMessage("Tra tra")
	m.print()
}
