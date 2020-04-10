package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person1 := new(Person) // person1 is a pointer to an instance of person
	person1.Name = "Shajal"
	person1.Age = 13
	fmt.Println(person1)
	fmt.Println(person1.Name)
	fmt.Println(person1.Age)
}
