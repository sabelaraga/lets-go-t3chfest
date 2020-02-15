//TODO: Add a Constructor
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

func NewMessage(/* What goes here?*/) /*And here?*/ {
	// TODO implement
}

func main() {
	// TODO change the next line to call the constructor
	m := Message{ txt : "Tra tra" }
	m.print()
}
