package main

import (
	"fmt"
)

type Message struct{
  txt string
}

func main() {
	m := Message{ txt : "Tra tra" }
	fmt.Println(m.txt)
}
