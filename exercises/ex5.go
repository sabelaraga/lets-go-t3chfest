//TODO: print returns an error if the message is empty
package main

import (
	"fmt"
)

type MyError struct{
	msg string
}

func (e MyError) Error() string{
	return e.msg
}


type Message struct{
  txt string
}

func (m Message) print () // return something here {
	// Add a check to see if the message is empty
	// (and return an error if it is)
	fmt.Println(m.txt)
	// If there is no error, return nil
}

func NewMessage(s string) Message {
	return Message{ txt : s }
}

func main() {
	m := NewMessage("Tra tra")
	// Check the return value and print an error message if there is an error
	m.print()
}
