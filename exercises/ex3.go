//TODO: Add Message as a receiver of print
package main

import (
	"fmt"
)

type Message struct{
  txt string
}

// Change this line so it has a receiver
func print (m Message) {
	fmt.Println(m.txt)
}

func main() {
	m := Message{ txt : "Tra tra" }
	// How do we call it?
	print (m)
	
}
