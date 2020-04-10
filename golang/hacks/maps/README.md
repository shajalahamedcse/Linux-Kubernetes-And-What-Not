# Maps

In golang, **map** is a data structure that provides you an unordered collection of **key/value** pairs. **Maps** are used to look up a value by its associated key. You store values into the map based on a key.

Map is implemented using a hash table.**Hash tables** in general are O(1) for inserting, looking up and deleting data.So we can assume that this is also true for Go's maps.In maps we can easily retrieve a value by providing the key. As maps are unordered collections so there's no way to predict the order in which the **key/value** pairs will be returned. In every iteration over a map could return data in a different order.

## Map initialization

Let's see how we can create a map in **golang**.

    package main
 
    import "fmt"
    
    var person = map[string]int{"Shajal": 10, "Sajib": 20}
    
    func main() {
        fmt.Println(person)
    }

## Empty Map Declaration

    package main
    
    import "fmt"
    
    func main() {
        var person = map[string]int{}
        fmt.Println(person)        // map[]
        fmt.Printf("%T\n", person) // map[string]int
    }

## Map Declaration Using make() Function

**make()** function takes the type of the **map** as argument and it returns an initialized map.

    package main
    
    import "fmt"
    
    func main() {
        var person = make(map[string]int)
        person["Shajal"] = 10
        person["Sajib"] = 20
        fmt.Println(person)
    }

## Map Length

We can use built-in **len()** functionTo determine how many items (key-value pairs) a map has.

    package main
 
    import "fmt"
    
    func main() {
        var person = make(map[string]int)
        person["Shajal"] = 10
        person["Sajib"] = 20
    
        // Empty Map
        personList := make(map[string]int)
    
        fmt.Println(len(person))     // 2
        fmt.Println(len(personList)) // 0
    }

**len()** function will return zero for an uninitialized map.

## Accessing Items

We can easily access the items of a map by referring to its key name, inside square brackets.This is an 
O(1) operation.

    package main
    
    import "fmt"
    
    func main() {
        var person = map[string]int{"Shajal": 10, "Sajib": 20}
    
        fmt.Println(person["Shajal"])// 10
    }

## Adding Items

    package main
    
    import "fmt"
    
    func main() {
        var people = map[string]int{"Shajal": 10, "Sajib": 20}
        fmt.Println(people) // Initial Map
    
        people["Fahad"] = 30 // Add element
        people["Shahed"] = 40
    
        fmt.Println(people)
    }

## Update Values

Changed the "Shahed" to 50

    package main
    
    import "fmt"
    
    func main() {
        var person = map[string]int{"Shajal": 10, "Shahed": 20}
        fmt.Println(person) // Initial Map
    
        person["Shahed"] = 50 // Edit item
        fmt.Println(person)
    }


## Delete Items

The built-in **delete()** function deletes an item from a given map associated with the provided key.

    package main
    
    import "fmt"
    
    func main() {
        var person = make(map[string]int)
        person["Sajib"] = 10
        person["Shajal"] = 20
        person["Fahad"] = 30
        person["Shahed"] = 40
    
        fmt.Println(person)
    
        delete(person, "Sajib")
        fmt.Println(person)
    }

## Iterate over a Map

The **for…range** loop statement can be used to iterate over a map.

    package main
    
    import "fmt"
    
    func main() {
        var person = map[string]int{"Shahed": 10, "Sajib": 20,
            "Fahad": 30, "Roni": 40, "Shajal": 50}

        for key, element := range person {
            fmt.Println("Key:", key, "=>", "Element:", element)
        }
    }

Each iteration returns a key and its correlated element content.

## Truncate Map

    package main
    
    func main() {
        var person = map[string]int{"Shajal": 10, "Sajib": 20}

        person = make(map[string]int)
    }
