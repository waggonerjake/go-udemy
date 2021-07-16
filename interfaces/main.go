package main

import "fmt"

//Interface will associate any type that has all the functions listed within the interface
//and will promote the types to become the type of interface. This allows for
//a pseudo-generic way to call things or almost like a polymorphic way. Just
//a common wrapper to allow for code reuse. Good way to know if the function should
//be in an interface is to ask yourself, "Is this function dependent on the type?".
//If it isn't, then it should go into an interface.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eng := englishBot{}
	spa := spanishBot{}

	printGreeting(eng)
	printGreeting(spa)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
