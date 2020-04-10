# Struct Instantiation using new keyword

We can create an instance of a **struct** with the **new** keyword. After that it is possible to assign data values to the data fields using **dot notation**.

## Code

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

Here,One instance of the Person struct is instantiated, **person1** points to the address of the instantiated struct.

## How to run

    $ go run main.go
