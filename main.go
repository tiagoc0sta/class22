/*package main

import "fmt"

type Person struct {
    name string
    age  int
}

// Getter method for name
func (p *Person) GetName() string {
    return p.name
}

// Setter method for name
func (p *Person) SetName(name string) {
    p.name = name
}

// Getter method for age
func (p *Person) GetAge() int {
    return p.age
}

// Setter method for age
func (p *Person) SetAge(age int) {
    p.age = age
}

func main() {
    // Create a new instance of Person
    person1 := Person{name: "Alice", age: 30}

    // Use getter methods to retrieve name and age
    fmt.Println("Name:", person1.GetName())
    fmt.Println("Age:", person1.GetAge())

    // Use setter methods to update name and age
    person1.SetName("Bob")
    person1.SetAge(25)
		fmt.Println()

    fmt.Println("Updated Name:", person1.GetName())
    fmt.Println("Updated Age:", person1.GetAge())

    // Create another instance of Person
    person2 := Person{name: "Charlie", age: 40}

    fmt.Println("Name:", person2.GetName())
    fmt.Println("Age:", person2.GetAge())
}
*/

package main

import "fmt"

type IceCream struct {
    flavor   string
    scoops   int
    toppings []string
}

// Getter method for flavor
func (i *IceCream) GetFlavor() string {
    return i.flavor
}

// Setter method for flavor
func (i *IceCream) SetFlavor(flavor string) {
    i.flavor = flavor
}

// Getter method for scoops
func (i *IceCream) GetScoops() int {
    return i.scoops
}

// Setter method for scoops
func (i *IceCream) SetScoops(scoops int) {
    i.scoops = scoops
}

// Getter method for toppings
func (i *IceCream) GetToppings() []string {
    return i.toppings
}

// Setter method for toppings
func (i *IceCream) SetToppings(toppings []string) {
    i.toppings = toppings
}

func main() {
    // Create a new instance of IceCream
    iceCream1 := IceCream{flavor: "Chocolate", scoops: 2, toppings: []string{"Caramel", "Coconut"}}

    // Use getter methods to retrieve flavor, scoops, and toppings
    fmt.Println("Flavor:", iceCream1.GetFlavor())
    fmt.Println("Scoops:", iceCream1.GetScoops())
    fmt.Println("Toppings:", iceCream1.GetToppings())

    // Use setter methods to update flavor, scoops, and toppings
    iceCream1.SetFlavor("Vanilla")
    iceCream1.SetScoops(3)
    iceCream1.SetToppings([]string{"Cream", "Cherry"})

    fmt.Println("Updated Flavor:", iceCream1.GetFlavor())
    fmt.Println("Updated Scoops:", iceCream1.GetScoops())
    fmt.Println("Updated Toppings:", iceCream1.GetToppings())

    // Create another instance of IceCream
    iceCream2 := IceCream{flavor: "Strawberry", scoops: 1, toppings: []string{"Nuts", "Caramel sauce"}}

    fmt.Println("Flavor:", iceCream2.GetFlavor())
    fmt.Println("Scoops:", iceCream2.GetScoops())
    fmt.Println("Toppings:", iceCream2.GetToppings())
}
