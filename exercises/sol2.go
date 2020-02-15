package main

import (
	"fmt"
)

type Message struct{
  txt string
}

func print (m Message) {
	fmt.Println(m.txt)
}

func main() {
	m := Message{ txt : "Tra tra" }
	print (m)
	
}
