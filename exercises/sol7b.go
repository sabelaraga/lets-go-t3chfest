package main

import (
	"fmt"
	"time"
	"sync"
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
        time.Sleep(1)
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
  input = []string {"Pienso en tu mirá", "tu mirá", "clavá", "una bala en el pecho", "Pienso en tu mirá"}
)

func makeList (iList []string, msgch chan Message ) {
  	for i, v := range iList {
	    isLast := (i == (len(iList)-1))
    		msgch <- NewMessage(v, i, isLast)
  	}
}

func printList (msgch chan Message){
  	for m := range msgch { 
		if err := m.print(); err!= nil {
			fmt.Printf("Ouch! %v", err)
	 	}
		if (m.isLast) {
			break
		}
  	}
	wait.Done()
}

var wait sync.WaitGroup

func main() {
	msgch := make (chan Message)
	wait.Add(1)
	go printList (msgch)
	makeList(input, msgch)
	wait.Wait()
}
