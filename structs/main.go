package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	me := person{
		firstName: "Jake",
		lastName:  "Waggoner",
		contactInfo: contactInfo{
			email: "waggonerjake99@gmail.com",
			zip:   46032,
		},
	}

	me.updateFirstName("Jacob")
	me.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//Using a * will let you pass by reference when using a receiver on a non-basic type
func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}
