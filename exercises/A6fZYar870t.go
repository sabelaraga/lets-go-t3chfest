// TODO: use chanels instead of lists to send messages between the methods
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

func (m Message) print (i int) error {
	if m.txt == "" {
     		return MyError { msg: "fail!"}
   	}

	fmt.Printf("(%d) %s\n", i, m.txt)
	return nil	
}

func NewMessage(s string) Message {
	return Message{ txt : s }
}

var (
  input = []string {"Pienso en tu mirá", "tu mirá", "clavá"}
)

func makeList (iList [] string) []Message{
  	oList := []Message{}
  	for _, v := range input {
    		oList = append(oList, NewMessage(v))
  	}
  	return oList
}

func printList (list []Message){
  	for i, m := range list {
		if err := m.print(i); err!= nil {
			fmt.Printf("Ouch! %v", err)
	 	}
  	}
}


func main() {
	list := makeList(input)
	printList(list)
}
