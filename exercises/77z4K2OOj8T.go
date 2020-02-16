// TODO: there is a problem here...
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
  index int
  isLast bool
}

func (m Message) print () error {
	if m.txt == "" {
     		return MyError { msg: "fail!"}
   	}

	fmt.Printf("(%d) %s\n", m.index, m.txt)
	return nil	
}

func NewMessage(s string, i int, isLast bool) Message {
	return Message{ txt : s, index: i, isLast: isLast }
}

var (
  input = []string {"Pienso en tu mirá", "tu mirá", "clavá"}
)

func makeList (iList []string, msgch chan Message ) {
	// put the new messages on a channel
  	for i, v := range iList {
	    isLast := (i == (len(iList)-1))
    		msgch <- NewMessage(v, i, isLast)
  	}
}

func printList (msgch chan Message){
	// Go through the channel instead of the list
	// Hint: Stop when you find the last message.
  	for m := range msgch { 
		if err := m.print(); err!= nil {
			fmt.Printf("Ouch! %v", err)
	 	}
		if (m.isLast) {
			break
		}
  	}
}



func main() {
	// 1. Create a channel
	msgch := make (chan Message)
	// 2. Launch printList as a goroutine and get a channel as a parameter
	go printList (msgch)
	// 3. Use the channel as a parameter to make the list
	makeList(input, msgch)
}
