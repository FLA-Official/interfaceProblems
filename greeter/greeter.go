package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p *Person) Greet() string {
	return ("Welcome " + p.Name)
}

type Robot struct {
	Name string
}

func (r *Robot) Greet() string {
	return ("New Bot named " + r.Name + " is added")
}

func main() {
	var greeter1 Greeter

	greeter1 = &Person{
		Name: "Farhan",
	}

	var greeter2 Greeter

	greeter2 = &Robot{
		Name: "Claude.ai",
	}

	fmt.Println(greeter1.Greet())
	fmt.Println(greeter2.Greet())
}
