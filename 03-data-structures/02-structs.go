package main

import "fmt"

type Person struct {
	age      int
	lastName string
	name     string
}

type SecretAgent struct {
	person Person
	lpm    bool
}

func main() {
	person := Person{name: "Iván", lastName: "Glez"}
	person.age = 5
	agent := SecretAgent{person, false}
	fmt.Println(agent.person)
	fmt.Println(person.name, person.lastName)
	fmt.Println(agent)
}
