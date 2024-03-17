/*
par convention il s'appelle main
*/
package main

import (
	"fmt"
	"log"
)

/*
doit s'appeler main
ne prends pas de parameters
entry point for the app
*/
func maind() {
	var name = "Ayoub"
	var whatToSay = saySomething(name)
	fmt.Println(whatToSay)

	var someThingToDo = "manger"
	var whatToDo, _ = doSomething(someThingToDo)
	log.Println(whatToDo)

	var myString string
	myString = "GREEN"
	log.Println("This is my string", myString)

	changeUsingPointer(&myString)
	log.Println("This is my string, after function call", myString)

}

func saySomething(s string) string {
	return s
}

// retourne 2 valeurs (utilser _ après pour stocker)
func doSomething(s string) (string, string) {
	return s, "world"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", *s)
	newValue := "Red"
	*s = newValue
}
