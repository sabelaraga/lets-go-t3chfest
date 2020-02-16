// TODO: use chanels instead of lists to send messages between the methods
// See hints in the main function, and, then, in printList and makeList
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

func makeList (iList [] string) []Message{
        // put the new messages on a channel
  	oList := []Message{}
  	for i, v := range input {
	        isLast := (i == (len(oList)-1))
    		oList = append(oList, NewMessage(v, i, isLast))
  	}
  	return oList
}

func printList (list []Message){
	// Go through the channel instead of the list
	// Hint: Stop when you find the last message.
  	for _, m := range list { 
		if err := m.print(); err!= nil {
			fmt.Printf("Ouch! %v", err)
	 	}
  	}
}


func main() {
	// 1. Create a channel
	// 2. Launch printList as a goroutine and get a channel as a parameter
	// 3. Use the channel as a parameter to make the list
	list := makeList(input)
	printList(list)
}
