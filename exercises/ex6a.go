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

func (m Message) print () error {
	if m.txt == "" {
     		return MyError { msg: "fail!"}
   	}

	fmt.Println(m.txt)
	return nil	
}

func NewMessage(s string) Message {
	return Message{ txt : s }
}

var (
  input = []string {"Pienso en tu mirá", "tu mirá", "clavá"}
)

func main() {
	// Go through the messages in the input and print them
	m := NewMessage("Tra tra")
	// 
	if err := m.print(); err!= nil {
		fmt.Printf("Ouch! %v", err)
	}
}
