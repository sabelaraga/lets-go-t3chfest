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

func main() {
	m := Message{ txt : "Tra tra" }
	m.print()
}

