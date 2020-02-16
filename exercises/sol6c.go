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

func makeList (iList [] string) []Message{
  	oList := []Message{}
  	for _, v := range iList {
    		oList = append(oList, NewMessage(v))
  	}
  	return oList
}

func printList (list []Message){
  	for _, m := range list {
		if err := m.print(); err!= nil {
			fmt.Printf("Ouch! %t", err)
	 	}
  	}
}


func main() {
	list := makeList(input)
	printList(list)
}
